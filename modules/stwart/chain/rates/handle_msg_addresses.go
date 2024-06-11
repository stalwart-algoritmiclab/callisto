/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

// handleMsgCreateAddresses allows to properly handle a MsgCreateAddresses
func (m *Module) handleMsgCreateAddresses(tx *juno.Tx, msg *rates.MsgCreateAddresses) error {
	return m.ratesRepo.InsertMsgCreateAddresses(tx.TxHash, &rates.MsgCreateAddresses{
		Creator: msg.Creator,
		Address: msg.Address,
	})
}

// handleMsgUpdateAddresses allows to properly handle a MsgUpdateAddresses
func (m *Module) handleMsgUpdateAddresses(tx *juno.Tx, msg *rates.MsgUpdateAddresses) error {
	return m.ratesRepo.InsertMsgUpdateAddresses(tx.TxHash, &rates.MsgUpdateAddresses{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	})
}

// handleMsgDeleteAddresses allows to properly handle a MsgDeleteAddresses
func (m *Module) handleMsgDeleteAddresses(tx *juno.Tx, msg *rates.MsgDeleteAddresses) error {
	return m.ratesRepo.InsertMsgDeleteAddresses(tx.TxHash, &rates.MsgDeleteAddresses{
		Creator: msg.Creator,
		Id:      msg.Id,
	})
}
