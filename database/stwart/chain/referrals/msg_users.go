/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
)

// InsertMsgCreateUser - insert a new MsgCreateUser into the database (stwart_referrals_create_user).
func (r Repository) InsertMsgCreateUser(hash string, msgs ...*types.MsgCreateUser) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_referrals_create_user (
			 creator, account_address, referrer, referrals, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id, creator, account_address, referrer, referrals, tx_hash
	`

	for _, msg := range msgs {
		m := toMsgCreateUserDatabase(hash, msg)
		if _, err := r.db.Exec(q, m.Creator, m.AccountAddress, m.Referrer, msg.Referrals, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// InsertMsgUpdateUser - insert a new MsgUpdateUser into the database (stwart_referrals_update_user).
func (r Repository) InsertMsgUpdateUser(hash string, msgs ...*types.MsgUpdateUser) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_referrals_update_user (
			 creator, account_address, referrer, referrals, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING id, creator, account_address, referrer, referrals, tx_hash
	`

	for _, msg := range msgs {
		m := toMsgUpdateUserDatabase(hash, msg)
		if _, err := r.db.Exec(q, m.Creator, m.AccountAddress, m.Referrer, msg.Referrals, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// InsertMsgDeleteUser - insert a new MsgDeleteUser into the database (stwart_referrals_delete_user).
func (r Repository) InsertMsgDeleteUser(hash string, msgs ...*types.MsgDeleteUser) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_referrals_delete_user (
			 creator, account_address, tx_hash
		) VALUES (
			$1, $2, $3
		) RETURNING id, creator, account_address, tx_hash
	`

	for _, msg := range msgs {
		m := toMsgDeleteUserDatabase(hash, msg)
		if _, err := r.db.Exec(q, m.Creator, m.AccountAddress, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
