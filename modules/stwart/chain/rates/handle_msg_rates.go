/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

// handleMsgCreateRates allows to properly handle a MsgCreateRates
func (m *Module) handleMsgCreateRates(tx *juno.Transaction, msg *types.MsgCreateRates) error {

	return m.ratesRepo.InsertMsgCreateRates(tx.TxHash, &types.MsgCreateRates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     msg.Rate,
		Decimals: msg.Decimals,
	})
}

// handleMsgUpdateRates allows to properly handle a MsgUpdateRates
func (m *Module) handleMsgUpdateRates(tx *juno.Transaction, msg *types.MsgUpdateRates) error {
	return m.ratesRepo.InsertMsgUpdateRates(tx.TxHash, &types.MsgUpdateRates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     msg.Rate,
		Decimals: msg.Decimals,
	})
}

func (m *Module) handleMsgDeleteRates(tx *juno.Transaction, msg *types.MsgDeleteRates) error {
	return m.ratesRepo.InsertMsgDeleteRates(tx.TxHash, &types.MsgDeleteRates{
		Creator: msg.Creator,
		Denom:   msg.Denom,
	})
}
