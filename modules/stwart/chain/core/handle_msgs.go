/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	juno "github.com/forbole/juno/v6/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
)

// handleMsgIssue handles the MsgIssue message
func (m *Module) handleMsgIssue(tx *juno.Transaction, _ int, msg *types.MsgIssue) error {
	return m.coreRepo.InsertMsgIssue(tx.TxHash, &types.MsgIssue{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Denom:   msg.Denom,
		Address: msg.Address,
	})
}

// handleMsgWithdraw handles the MsgWithdraw message
func (m *Module) handleMsgWithdraw(tx *juno.Transaction, _ int, msg *types.MsgWithdraw) error {
	return m.coreRepo.InsertMsgWithdraw(tx.TxHash, &types.MsgWithdraw{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Denom:   msg.Denom,
		Address: msg.Address,
	})
}

// handleMsgSend handles the MsgSend message
func (m *Module) handleMsgSend(tx *juno.Transaction, _ int, msg *types.MsgSend) error {
	return m.coreRepo.InsertMsgSend(tx.TxHash, &types.MsgSend{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		From:    msg.From,
		To:      msg.To,
		Denom:   msg.Denom,
	})
}
