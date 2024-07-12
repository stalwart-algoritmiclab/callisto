/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	juno "github.com/forbole/juno/v6/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

// msgFilter defines the messages that should be handled by this module
var msgFilter = map[string]bool{
	"/stwartchain.core.MsgWithdraw":  true,
	"/stwartchain.core.MsgIssue":     true,
	"/stwartchain.core.MsgRefund":    true,
	"/stwartchain.core.MsgRefReward": true,
	"/stwartchain.core.MsgSend":      true,
	"/stwartchain.core.MsgFees":      true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.core.MsgWithdraw":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgWithdraw{})
		return m.handleMsgWithdraw(tx, index, cosmosMsg)
	case "/stwartchain.core.MsgIssue":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgIssue{})
		return m.handleMsgIssue(tx, index, cosmosMsg)
	case "/stwartchain.core.MsgRefund":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgRefund{})
		return m.handleMsgRefund(tx, index, cosmosMsg)
	case "/stwartchain.core.MsgRefReward":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgRefReward{})
		return m.handleMsgRefReward(tx, index, cosmosMsg)
	case "/stwartchain.core.MsgSend":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgSend{})
		return m.handleMsgSend(tx, index, cosmosMsg)
	case "/stwartchain.core.MsgFees":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgFees{})
		return m.handleMsgFees(tx, index, cosmosMsg)

	default:
		return nil
	}
}
