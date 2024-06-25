/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
)

// handleMsgCreateAddresses allows to properly handle a MsgCreateAddresses
func (m *Module) handleMsgCreateAddresses(tx *juno.Transaction, _ int, msg *types.MsgCreateAddresses) error {
	return m.feepolicyRepo.InsertMsgCreateAddresses(tx.TxHash, msg)
}

// handleMsgUpdateAddresses allows to properly handle a MsgUpdateAddresses
func (m *Module) handleMsgUpdateAddresses(tx *juno.Transaction, _ int, msg *types.MsgUpdateAddresses) error {
	return m.feepolicyRepo.InsertMsgUpdateAddresses(tx.TxHash, msg)
}

// handleMsgDeleteAddresses allows to properly handle a MsgDeleteAddresses
func (m *Module) handleMsgDeleteAddresses(tx *juno.Transaction, _ int, msg *types.MsgDeleteAddresses) error {
	return m.feepolicyRepo.InsertMsgDeleteAddresses(tx.TxHash, msg)
}
