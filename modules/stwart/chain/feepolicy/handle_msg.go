/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

var msgFilter = map[string]bool{
	"/stwartchain.feepolicy.MsgCreateTariffs":   true,
	"/stwartchain.feepolicy.MsgUpdateTariffs":   true,
	"/stwartchain.feepolicy.MsgDeleteTariffs":   true,
	"/stwartchain.feepolicy.MsgCreateAddresses": true,
	"/stwartchain.feepolicy.MsgUpdateAddresses": true,
	"/stwartchain.feepolicy.MsgDeleteAddresses": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.feepolicy.MsgCreateTariffs":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreateTariffs{})
		return m.handleMsgCreateTariffs(tx, index, cosmosMsg)
	case "/stwartchain.feepolicy.MsgUpdateTariffs":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgUpdateTariffs{})
		return m.handleMsgUpdateTariffs(tx, index, cosmosMsg)
	case "/stwartchain.feepolicy.MsgDeleteTariffs":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgDeleteTariffs{})
		return m.handleMsgDeleteTariffs(tx, index, cosmosMsg)
	case "/stwartchain.feepolicy.MsgCreateAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreateAddresses{})
		return m.handleMsgCreateAddresses(tx, index, cosmosMsg)
	case "/stwartchain.feepolicy.MsgUpdateAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgUpdateAddresses{})
		return m.handleMsgUpdateAddresses(tx, index, cosmosMsg)
	case "/stwartchain.feepolicy.MsgDeleteAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgDeleteAddresses{})
		return m.handleMsgDeleteAddresses(tx, index, cosmosMsg)
	default:
		return nil
	}
}
