/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package stwart

import (
	"encoding/json"
	"fmt"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v6/modules"
	"github.com/forbole/juno/v6/types"
	txtypes "github.com/forbole/juno/v6/types"

	dbhelpers "github.com/stalwart-algoritmiclab/callisto/database/types"
	dbtypes "github.com/stalwart-algoritmiclab/callisto/database/types"
)

const (
	module = "stwart"
)

// parseMissingBlocksAndTransactions - parse missing blocks and transactions
func (m *Module) parseMissingBlocksAndTransactions(height int64) (dbtypes.BlockRow, []*txtypes.Transaction, error) {
	block, err := m.node.Block(height)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, fmt.Errorf("failed to get block from node: %s", err)
	}

	events, err := m.node.BlockResults(height)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, fmt.Errorf("failed to get block results from node: %s", err)
	}

	txs, err := m.node.Txs(block)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, fmt.Errorf("failed to get transactions for block: %s", err)
	}

	vals, err := m.node.Validators(height)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, fmt.Errorf("failed to get validators for block: %s", err)
	}

	return m.ExportBlock(block, events, txs, vals)
}

// ExportBlock accepts a finalized block and a corresponding set of transactions
// and persists them to the database along with attributable metadata. An error
// is returned if the writing fails.
func (m *Module) ExportBlock(
	b *tmctypes.ResultBlock, r *tmctypes.ResultBlockResults, txs []*txtypes.Transaction, vals *tmctypes.ResultValidators,
) (dbtypes.BlockRow, []*txtypes.Transaction, error) {
	// Save all validators
	err := m.SaveValidators(vals.Validators)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, err
	}

	// Make sure the proposer exists
	proposerAddr := sdk.ConsAddress(b.Block.ProposerAddress).String()
	val := findValidatorByAddr(proposerAddr, vals)
	if val == nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{},
			fmt.Errorf("failed to find validator by proposer address %s: %s", proposerAddr, err)
	}

	block := types.NewBlockFromTmBlock(b, sumGasTxs(txs))

	// Save the block
	err = m.db.SaveBlock(block)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, fmt.Errorf("failed to persist block: %s", err)
	}

	// Save the commits
	err = m.ExportCommit(b.Block.LastCommit, vals)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, err
	}

	// Call the block handlers
	for _, mod := range m.stwartModules {
		if blockModule, ok := mod.(modules.BlockModule); ok {
			err = blockModule.HandleBlock(b, r, txs, vals)
			if err != nil {
				m.logger.BlockError(mod, b, err)
			}
		}
	}

	tsx, err := m.ExportTxs(txs)
	if err != nil {
		return dbtypes.BlockRow{}, []*txtypes.Transaction{}, err
	}

	// Export the transactions
	return dbtypes.BlockRow{
		Height:          block.Height,
		Hash:            block.Hash,
		TxNum:           int64(block.TxNum),
		TotalGas:        int64(block.TotalGas),
		ProposerAddress: dbhelpers.ToNullString(block.ProposerAddress),
		Timestamp:       block.Timestamp,
	}, tsx, nil
}

// SaveValidators persists a list of Tendermint validators with an address and a
// consensus public key. An error is returned if the public key cannot be Bech32
// encoded or if the DB write fails.
func (m *Module) SaveValidators(vals []*tmtypes.Validator) error {
	var validators = make([]*types.Validator, len(vals))
	for index, val := range vals {
		consAddr := sdk.ConsAddress(val.Address).String()

		consPubKey, err := types.ConvertValidatorPubKeyToBech32String(val.PubKey)
		if err != nil {
			return fmt.Errorf("failed to convert validator public key for validators %s: %s", consAddr, err)
		}

		validators[index] = types.NewValidator(consAddr, consPubKey)
	}

	err := m.db.SaveValidators(validators)
	if err != nil {
		return fmt.Errorf("error while saving validators: %s", err)
	}

	return nil
}

// findValidatorByAddr finds a validator by a consensus address given a set of
// Tendermint validators for a particular block. If no validator is found, nil
// is returned.
func findValidatorByAddr(consAddr string, vals *tmctypes.ResultValidators) *tmtypes.Validator {
	for _, val := range vals.Validators {
		if consAddr == sdk.ConsAddress(val.Address).String() {
			return val
		}
	}

	return nil
}

// sumGasTxs returns the total gas consumed by a set of transactions.
func sumGasTxs(txs []*types.Transaction) uint64 {
	var totalGas uint64

	for _, tx := range txs {
		totalGas += tx.GasUsed
	}

	return totalGas
}

// ExportCommit accepts a block commitment and a corresponding set of
// validators for the commitment and persists them to the database. An error is
// returned if any writing fails or if there is any missing-aggregated data.
func (m *Module) ExportCommit(commit *tmtypes.Commit, vals *tmctypes.ResultValidators) error {
	var signatures []*types.CommitSig
	for _, commitSig := range commit.Signatures {
		// Avoid empty commits
		if commitSig.Signature == nil {
			continue
		}

		valAddr := sdk.ConsAddress(commitSig.ValidatorAddress)
		val := findValidatorByAddr(valAddr.String(), vals)
		if val == nil {
			return fmt.Errorf("failed to find validator by commit validator address %s", valAddr.String())
		}

		signatures = append(signatures, types.NewCommitSig(
			types.ConvertValidatorAddressToBech32String(commitSig.ValidatorAddress),
			val.VotingPower,
			val.ProposerPriority,
			commit.Height,
			commitSig.Timestamp,
		))
	}

	err := m.db.SaveCommitSignatures(signatures)
	if err != nil {
		return fmt.Errorf("error while saving commit signatures: %s", err)
	}

	return nil
}

// ExportTxs accepts a slice of transactions and persists then inside the database.
// An error is returned if the write fails.
func (m *Module) ExportTxs(txs []*types.Transaction) ([]*txtypes.Transaction, error) {
	// handle all transactions inside the block
	for _, tx := range txs {
		events := make([]sdk.Event, 0)

		for _, ev := range tx.Events {
			events = append(events, sdk.Event{Type: ev.Type, Attributes: ev.Attributes})
		}

		msgLog := sdk.NewABCIMessageLog(0, "", events)
		tx.Logs = sdk.ABCIMessageLogs{msgLog}

		// save the transaction
		err := m.saveTx(tx)
		if err != nil {
			return []*txtypes.Transaction{}, fmt.Errorf("error while storing txs: %s", err)
		}

		// call the tx handlers
		m.handleTx(tx)

		// call the msg handlers
		for i, msg := range tx.Tx.Body.Messages {
			m.handleMessage(i, msg, tx)
		}
	}

	return txs, nil
}

// handleTx accepts the transaction and calls the tx handlers.
func (m *Module) handleTx(tx *types.Transaction) {
	// Call the tx handlers
	for _, module := range m.stwartModules {
		if transactionModule, ok := module.(modules.TransactionModule); ok {
			err := transactionModule.HandleTx(tx)
			if err != nil {
				m.logger.TxError(module, tx, err)
			}
		}
	}
}

// handleMessage accepts the transaction and handles messages contained
// inside the transaction.
func (m *Module) handleMessage(index int, msg types.Message, tx *types.Transaction) {
	// Allow modules to handle the message
	for _, module := range m.stwartModules {
		if messageModule, ok := module.(modules.MessageModule); ok {
			err := messageModule.HandleMsg(index, msg, tx)
			if err != nil {
				m.logger.MsgError(module, tx, msg, err)
			}
		}

		// If it's a MsgExecute, we need to make sure the included messages are handled as well
		if msg.GetType() == "/cosmos.authz.v1beta1.MsgExec" {
			var msgExec struct {
				Msgs []json.RawMessage `json:"msgs"`
			}

			err := json.Unmarshal(msg.GetBytes(), &msgExec)
			if err != nil {
				m.logger.Error("unable to unmarshal MsgExec inner messages", "error", err)
				return
			}

			for authzIndex, msgAny := range msgExec.Msgs {
				executedMsg, err := types.UnmarshalMessage(authzIndex, msgAny)
				if err != nil {
					m.logger.Error("unable to unpack MsgExec inner message", "index", authzIndex, "error", err)
				}

				for _, module := range m.stwartModules {
					if messageModule, ok := module.(modules.AuthzMessageModule); ok {
						err = messageModule.HandleMsgExec(index, authzIndex, executedMsg, tx)
						if err != nil {
							m.logger.MsgError(module, tx, executedMsg, err)
						}
					}
				}
			}
		}
	}
}

// saveTx accepts the transaction and persists it inside the database.
// An error is returned if the write fails.
func (m *Module) saveTx(tx *types.Transaction) error {
	err := m.db.SaveTx(tx)
	if err != nil {
		return fmt.Errorf("failed to handle transaction with hash %s: %s", tx.TxResponse.TxHash, err)
	}
	return nil
}
