/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/core"
)

// handleMsgIssue handles the MsgIssue message
func (m *Module) handleMsgIssue(tx *juno.Tx, _ int, msg *core.MsgIssue) error {
	return m.coreRepo.InsertMsgIssue(tx.TxHash, &core.MsgIssue{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Denom:   msg.Denom,
		Address: msg.Address,
	})
}

func (m *Module) handleMsgWithdraw(tx *juno.Tx, _ int, msg *core.MsgWithdraw) error {
	return m.coreRepo.InsertMsgWithdraw(tx.TxHash, &core.MsgWithdraw{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		Denom:   msg.Denom,
		Address: msg.Address,
	})
}

func (m *Module) handleMsgFees(tx *juno.Tx, _ int, msg *core.MsgFees) error {
	return m.coreRepo.InsertMsgFees(tx.TxHash, &core.MsgFees{
		Creator:   msg.Creator,
		Comission: msg.Comission, //nolint:misspell
		AddressTo: msg.AddressTo,
	})
}

func (m *Module) handleMsgRefund(tx *juno.Tx, _ int, msg *core.MsgRefund) error {
	return m.coreRepo.InsertMsgRefund(tx.TxHash, &core.MsgRefund{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		From:    msg.From,
		To:      msg.To,
	})
}

func (m *Module) handleMsgRefReward(tx *juno.Tx, _ int, msg *core.MsgRefReward) error {
	return m.coreRepo.InsertMsgRefReward(tx.TxHash, &core.MsgRefReward{
		Creator:  msg.Creator,
		Amount:   msg.Amount,
		Referrer: msg.Referrer,
	})
}

func (m *Module) handleMsgSend(tx *juno.Tx, _ int, msg *core.MsgSend) error {
	return m.coreRepo.InsertMsgSend(tx.TxHash, &core.MsgSend{
		Creator: msg.Creator,
		Amount:  msg.Amount,
		From:    msg.From,
		To:      msg.To,
		Denom:   msg.Denom,
	})
}
