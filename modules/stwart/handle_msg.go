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
	"sync"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/types"
	junoconf "github.com/forbole/juno/v5/types/config"

	dbtypes "github.com/stalwart-algoritmiclab/callisto/database/types"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

const intervalLastBlock = 2 * time.Second

// scheduler runs the scheduler
func (m *Module) scheduler() {
	ticker := time.NewTicker(intervalLastBlock)
	defer ticker.Stop()

	for range ticker.C {
		// get the latest-parsed block from a database
		lastRepoBlock, err := m.lastBlockRepo.Get()
		if err != nil {
			m.logger.Error("Fail lastBlockRepo.Get", "module", m.Name(), "error", err)
			os.Exit(1)
		}

		// get the latest block from node
		lastNodeBlockInt, err := m.node.LatestHeight()
		if err != nil {
			m.logger.Error("Fail node.LatestHeight", "module", m.Name(), "error", err)
			os.Exit(1)
		}

		lastNodeBlock := uint64(lastNodeBlockInt)

		// compare the latest block in the database and the latest block in the node
		if lastRepoBlock >= lastNodeBlock {
			continue
		}

		wg := &sync.WaitGroup{}
		blockChan := make(chan uint64, lastNodeBlock-lastRepoBlock)

		numWorkers := int(junoconf.Cfg.Parser.Workers)
		for i := 0; i < numWorkers; i++ {
			wg.Add(1)

			// worker function to process blocks
			go func() {
				defer wg.Done()
				for height := range blockChan {
					if err = m.parseBlock(height); err != nil {
						if errors.As(err, &errs.NotFound{}) {
							continue
						}

						m.logger.Error("Fail parseBlock", "module", m.Name(), "error", err)
						continue
					}

					if err = m.lastBlockRepo.Update(height); err != nil {
						m.logger.Error("Fail lastBlockRepo.Update", "module", m.Name(), "error", err)
						os.Exit(1)
					}
				}
			}()
		}

		// distribute blocks to workers
		for i := lastRepoBlock + 1; i <= lastNodeBlock; i++ {
			blockChan <- i
		}

		close(blockChan)
		wg.Wait()
	}
}

// parseBlock parse block
func (m *Module) parseBlock(lastBlock uint64) error {
	if _, _, err := m.parseMissingBlocksAndTransactions(int64(lastBlock)); err != nil {
		m.logger.Error("Fail parseMissingBlocksAndTransactions", "module", m.Name(), "error", err)
		return errs.Internal{Cause: "Fail parseMissingBlocksAndTransactions, error: " + err.Error()}
	}

	block, err := m.db.GetBlock(filter.NewFilter().SetArgument(dbtypes.FieldHeight, lastBlock))
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			if block, _, err = m.parseMissingBlocksAndTransactions(int64(lastBlock)); err != nil {
				m.logger.Error("Fail parseMissingBlocksAndTransactions", "module", m.Name(), "error", err)
				return errs.Internal{Cause: "Fail parseMissingBlocksAndTransactions, error: " + err.Error()}
			}
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
	txs, err := m.db.GetTransactions(filter.NewFilter().SetArgument(dbtypes.FieldHeight, block.Height))
	if err != nil {
		if errors.As(err, &errs.NotFound{}) {
			if block, _, err = m.parseMissingBlocksAndTransactions(block.Height); err != nil {
				m.logger.Error("Fail parseMissingBlocksAndTransactions", "module", m.Name(), "error", err)
				return errs.Internal{Cause: "Fail parseMissingBlocksAndTransactions, error: " + err.Error()}
			}
			return err
		}

		return errs.Internal{Cause: err.Error()}
	}

	if err = block.CheckTxNumCount(int64(len(txs))); err != nil {
		if _, txs, err = m.parseMissingBlocksAndTransactions(block.Height); err != nil {
			m.logger.Error("Fail parseMissingBlocksAndTransactions", "module", m.Name(), "error", err)
			return errs.Internal{Cause: "Fail parseMissingBlocksAndTransactions, error: " + err.Error()}
		}

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
