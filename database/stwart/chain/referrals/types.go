/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

const (
	tableMsgSetReferrer = "stwart_referrals_set_referrer"
)

type (
	// MsgSetReferrer - db model for 'stwart_referrals_set_referrer'
	MsgSetReferrer struct {
		ID       uint64 `db:"id"`
		Creator  string `db:"creator"`
		Referrer string `db:"referrer"`
		Referral string `db:"referrals"`
		TxHash   string `db:"tx_hash"`
	}
)
