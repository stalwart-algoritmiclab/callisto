/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import (
	juno "github.com/forbole/juno/v6/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

var msgFilter = map[string]bool{
	"/stwartchain.rates.MsgCreateRates":     true,
	"/stwartchain.rates.MsgUpdateRates":     true,
	"/stwartchain.rates.MsgDeleteRates":     true,
	"/stwartchain.rates.MsgCreateAddresses": true,
	"/stwartchain.rates.MsgUpdateAddresses": true,
	"/stwartchain.rates.MsgDeleteAddresses": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.rates.MsgCreateRates":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreateRates{})
		return m.handleMsgCreateRates(tx, cosmosMsg)
	case "/stwartchain.rates.MsgUpdateRates":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgUpdateRates{})
		return m.handleMsgUpdateRates(tx, cosmosMsg)
	case "/stwartchain.rates.MsgDeleteRates":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgDeleteRates{})
		return m.handleMsgDeleteRates(tx, cosmosMsg)
	case "/stwartchain.rates.MsgCreateAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreateAddresses{})
		return m.handleMsgCreateAddresses(tx, cosmosMsg)
	case "/stwartchain.rates.MsgUpdateAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgUpdateAddresses{})
		return m.handleMsgUpdateAddresses(tx, cosmosMsg)
	case "/stwartchain.rates.MsgDeleteAddresses":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgDeleteAddresses{})
		return m.handleMsgDeleteAddresses(tx, cosmosMsg)
	default:
		return nil
	}
}
