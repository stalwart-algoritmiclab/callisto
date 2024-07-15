/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"
)

// toMsgSetReferrerDatabase converts MsgSetReferrer to MsgSetReferrerDatabase
func toMsgSetReferrerDatabase(hash string, msg *types.MsgSetReferrer) MsgSetReferrer {
	return MsgSetReferrer{
		Creator:  msg.Creator,
		Referrer: msg.ReferrerAddress,
		Referral: msg.ReferrerAddress,
		TxHash:   hash,
	}
}
