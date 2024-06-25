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

// handleMsgCreateAddresses allows to properly handle a MsgCreateAddresses
func (m *Module) handleMsgCreateAddresses(tx *juno.Transaction, msg *types.MsgCreateAddresses) error {
	return m.ratesRepo.InsertMsgCreateAddresses(tx.TxHash, &types.MsgCreateAddresses{
		Creator: msg.Creator,
		Address: msg.Address,
	})
}

// handleMsgUpdateAddresses allows to properly handle a MsgUpdateAddresses
func (m *Module) handleMsgUpdateAddresses(tx *juno.Transaction, msg *types.MsgUpdateAddresses) error {
	return m.ratesRepo.InsertMsgUpdateAddresses(tx.TxHash, &types.MsgUpdateAddresses{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	})
}

// handleMsgDeleteAddresses allows to properly handle a MsgDeleteAddresses
func (m *Module) handleMsgDeleteAddresses(tx *juno.Transaction, msg *types.MsgDeleteAddresses) error {
	return m.ratesRepo.InsertMsgDeleteAddresses(tx.TxHash, &types.MsgDeleteAddresses{
		Creator: msg.Creator,
		Id:      msg.Id,
	})
}
