/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	juno "github.com/forbole/juno/v6/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"
)

// handleMsgSetReferrer handles a MsgSetReferrer message
func (m *Module) handleMsgSetReferrer(tx *juno.Transaction, msg *types.MsgSetReferrer) error {
	return m.referralsRepo.InsertMsgSetReferrer(tx.TxHash, msg)
}
