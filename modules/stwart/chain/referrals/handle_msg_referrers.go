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

// handleMsgSetReferrer handles a MsgSetReferrer message
func (m *Module) handleMsgSetReferrer(tx *juno.Tx, msg *referrals.MsgSetReferrer) error {
	return m.referralsRepo.InsertMsgSetReferrer(tx.TxHash, msg)
}
