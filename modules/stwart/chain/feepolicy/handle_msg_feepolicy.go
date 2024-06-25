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

// handleMsgCreateTariffs allows to properly handle a MsgCreateTariffs
func (m *Module) handleMsgCreateTariffs(tx *juno.Transaction, _ int, msg *types.MsgCreateTariffs) error {
	return m.feepolicyRepo.InsertMsgCreateTariffs(tx.Height, tx.TxHash, msg)
}

// handleMsgUpdateTariffs allows to properly handle a MsgUpdateTariffs
func (m *Module) handleMsgUpdateTariffs(tx *juno.Transaction, _ int, msg *types.MsgUpdateTariffs) error {
	return m.feepolicyRepo.InsertMsgUpdateTariffs(tx.Height, tx.TxHash, msg)
}

// handleMsgDeleteTariffs allows to properly handle a MsgDeleteTariffs
func (m *Module) handleMsgDeleteTariffs(tx *juno.Transaction, _ int, msg *types.MsgDeleteTariffs) error {
	return m.feepolicyRepo.InsertDeleteMsgDeleteTariffs(tx.Height, tx.TxHash, msg)
}
