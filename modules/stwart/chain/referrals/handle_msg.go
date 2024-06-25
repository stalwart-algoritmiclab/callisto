/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

var msgFilter = map[string]bool{
	"/stwartchain.referral.MsgCreateUser":  true,
	"/stwartchain.referral.MsgUpdateUser":  true,
	"/stwartchain.referral.MsgDeleteUser":  true,
	"/stwartchain.referral.MsgSetReferrer": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.referral.MsgCreateUser":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreateUser{})
		return m.handleMsgCreateUser(tx, cosmosMsg)
	case "/stwartchain.referral.MsgUpdateUser":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgUpdateUser{})
		return m.handleMsgUpdateUser(tx, cosmosMsg)
	case "/stwartchain.referral.MsgDeleteUser":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgDeleteUser{})
		return m.handleMsgDeleteUser(tx, cosmosMsg)
	case "/stwartchain.referral.MsgSetReferrer":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgSetReferrer{})
		return m.handleMsgSetReferrer(tx, cosmosMsg)

	default:
		return nil
	}
}
