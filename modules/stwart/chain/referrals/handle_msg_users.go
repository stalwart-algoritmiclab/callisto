/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/referrals"
)

// handleMsgCreateUser handles a MsgCreateUser message
func (m *Module) handleMsgCreateUser(tx *juno.Tx, msg *referrals.MsgCreateUser) error {
	return m.referralsRepo.InsertMsgCreateUser(tx.TxHash, msg)
}

// handleMsgUpdateUser handles a MsgUpdateUser message
func (m *Module) handleMsgUpdateUser(tx *juno.Tx, msg *referrals.MsgUpdateUser) error {
	return m.referralsRepo.InsertMsgUpdateUser(tx.TxHash, msg)
}

// handleMsgDeleteUser handles a MsgDeleteUser message
func (m *Module) handleMsgDeleteUser(tx *juno.Tx, msg *referrals.MsgDeleteUser) error {
	return m.referralsRepo.InsertMsgDeleteUser(tx.TxHash, msg)
}
