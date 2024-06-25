/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"
)

// handleMsgCreateUser handles a MsgCreateUser message
func (m *Module) handleMsgCreateUser(tx *juno.Transaction, msg *types.MsgCreateUser) error {
	return m.referralsRepo.InsertMsgCreateUser(tx.TxHash, msg)
}

// handleMsgUpdateUser handles a MsgUpdateUser message
func (m *Module) handleMsgUpdateUser(tx *juno.Transaction, msg *types.MsgUpdateUser) error {
	return m.referralsRepo.InsertMsgUpdateUser(tx.TxHash, msg)
}

// handleMsgDeleteUser handles a MsgDeleteUser message
func (m *Module) handleMsgDeleteUser(tx *juno.Transaction, msg *types.MsgDeleteUser) error {
	return m.referralsRepo.InsertMsgDeleteUser(tx.TxHash, msg)
}
