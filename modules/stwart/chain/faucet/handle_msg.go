/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package faucet

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

var msgFilter = map[string]bool{
	"/stwartchain.faucet.MsgIssue": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.faucet.MsgIssue":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgIssue{})
		return m.handleMsgIssue(tx, index, cosmosMsg)
	default:
		return nil
	}
}
