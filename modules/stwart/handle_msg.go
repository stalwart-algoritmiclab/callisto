/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package stwart

import (
	"errors"
	"fmt"
	"os"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/types"

	dbtypes "github.com/stalwart-algoritmiclab/callisto/database/types"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// scheduler runs the scheduler
func (m *Module) scheduler() {
	for {
		lastBlock, err := m.lastBlockRepo.Get()
		if err != nil {
			m.logger.Error("Fail lastBlockRepo.Get", "module", "stalwart", "error", err)
			continue
		}

		lastBlock++

		if err := m.parseBlock(lastBlock); err != nil {
			time.Sleep(time.Second)

			if errors.Is(err, errs.NotFound{}) {
				m.logger.Error("Fail parseBlock", "module", "stalwart", "error", err)
				continue
			}

			if _, _, err := m.parseMissingBlocksAndTransactions(int64(lastBlock)); err != nil {
				m.logger.Error("Fail parseMissingBlocksAndTransactions", "module", m.Name(), "error", err)
				continue
			}
		}

		if err = m.lastBlockRepo.Update(lastBlock); err != nil {
			m.logger.Error("Fail lastBlockRepo.Update", "module", "stalwart", "error", err)
			os.Exit(1)
		}
	}
}

// parseBlock parse block
func (m *Module) parseBlock(lastBlock uint64) error {
	block, err := m.db.GetBlock(filter.NewFilter().SetArgument(dbtypes.FieldHeight, lastBlock))
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			return err
		}

		return errs.Internal{Cause: err.Error()}
	}

	m.logger.Debug("parse block", "height", block.Height)

	if block.TxNum == 0 {
		return nil
	}

	return m.parseTx(block)
}

// parseTx parse txs from block
func (m *Module) parseTx(block dbtypes.BlockRow) error {
	if _, _, err := m.parseMissingBlocksAndTransactions(block.Height); err != nil {
		m.logger.Error("Fail parseMissingBlocksAndTransactions", "module", m.Name(), "error", err)
		return m.handleErrors(err)
	}

	txs, err := m.db.GetTransactions(filter.NewFilter().SetArgument(dbtypes.FieldHeight, block.Height))
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			return err
		}

		return errs.Internal{Cause: err.Error()}
	}

	if err = block.CheckTxNumCount(int64(len(txs))); err != nil {
		if err = block.CheckTxNumCount(int64(len(txs))); err != nil {
			return err
		}
	}

	for _, tx := range txs {
		if !tx.Successful() {
			continue
		}

		if err = m.parseMessages(tx); err != nil {
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// parseMessages - parse messages from transaction
func (m *Module) parseMessages(tx *types.Tx) error {
	for i, msg := range tx.Body.Messages {
		var stdMsg sdk.Msg
		if err := m.cdc.UnpackAny(msg, &stdMsg); err != nil {
			return fmt.Errorf("error while an unpacking message: %s", err)
		}

		for _, module := range m.stwartModules {
			if messageModule, ok := module.(modules.MessageModule); ok {
				if err := messageModule.HandleMsg(i, stdMsg, tx); err != nil {
					if errors.As(err, &errs.NotFound{}) {
						continue
					}

					m.logger.MsgError(module, tx, stdMsg, err)
					return err
				}
			}
		}
	}

	return nil
}
