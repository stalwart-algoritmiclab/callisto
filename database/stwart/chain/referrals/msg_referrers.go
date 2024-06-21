/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/referrals"
)

// InsertMsgSetReferrer - insert a new MsgSetReferrer into the database (stwart_referrals_set_referrer).
func (r Repository) InsertMsgSetReferrer(hash string, msgs ...*referrals.MsgSetReferrer) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_referrals_set_referrer (
			 creator, referrer, referral, tx_hash
		) VALUES (
			$1, $2, $3, $4
		) RETURNING id, creator, referrer, referral, tx_hash
	`

	for _, msg := range msgs {
		m := toMsgSetReferrerDatabase(hash, msg)
		if _, err := r.db.Exec(q, m.Creator, m.Referrer, m.Referral, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
