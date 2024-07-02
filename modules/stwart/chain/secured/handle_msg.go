/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

var msgFilter = map[string]bool{
	"/stwartchain.secured.MsgCreateAddresses": true,
	"/stwartchain.secured.MsgDeleteAddresses": true,
	"/stwartchain.secured.MsgUpdateAddresses": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.secured.MsgCreateAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreateAddresses{})
		return m.handleMsgCreateAddresses(tx, index, cosmosMsg)
	case "/stwartchain.secured.MsgDeleteAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgDeleteAddresses{})
		return m.handleMsgDeleteAddresses(tx, index, cosmosMsg)
	case "/stwartchain.secured.MsgUpdateAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgUpdateAddresses{})
		return m.handleMsgUpdateAddresses(tx, index, cosmosMsg)
	default:
		return nil
	}
}
