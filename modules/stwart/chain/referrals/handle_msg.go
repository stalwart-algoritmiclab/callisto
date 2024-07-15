/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	juno "github.com/forbole/juno/v6/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

// msgFilter defines the messages that should be handled by this module
var msgFilter = map[string]bool{
	"/stwartchain.referral.MsgSetReferrer": true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.referral.MsgSetReferrer":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgSetReferrer{})
		return m.handleMsgSetReferrer(tx, cosmosMsg)

	default:
		return nil
	}
}
