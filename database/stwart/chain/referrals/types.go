/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import "github.com/lib/pq"

const (
	// tableMsgCreateUser - table name for 'stwart_referrals_create_user'
	tableMsgCreateUser = "stwart_referrals_create_user"
	// tableMsgUpdateUser - table name for 'stwart_referrals_update_user'
	tableMsgUpdateUser = "stwart_referrals_update_user"
	// tableMsgDeleteUser - table name for 'stwart_referrals_delete_user'
	tableMsgDeleteUser = "stwart_referrals_delete_user"
	// tableMsgSetReferrer - table name for 'stwart_referrals_set_referrer'
	tableMsgSetReferrer = "stwart_referrals_set_referrer"
)

type (
	// MsgCreateUser - db model for 'stwart_referrals_create_user'
	MsgCreateUser struct {
		ID             uint64         `db:"id"`
		Creator        string         `db:"creator"`
		AccountAddress string         `db:"account_address"`
		Referrer       string         `db:"referrer"`
		Referrals      pq.StringArray `db:"referrals"`
		TxHash         string         `db:"tx_hash"`
	}

	// MsgUpdateUser - db model for 'stwart_referrals_update_user'
	MsgUpdateUser struct {
		ID             uint64         `db:"id"`
		Creator        string         `db:"creator"`
		AccountAddress string         `db:"account_address"`
		Referrer       string         `db:"referrer"`
		Referrals      pq.StringArray `db:"referrals"`
		TxHash         string         `db:"tx_hash"`
	}

	// MsgDeleteUser - db model for 'stwart_referrals_delete_user'
	MsgDeleteUser struct {
		ID             uint64 `db:"id"`
		Creator        string `db:"creator"`
		AccountAddress string `db:"account_address"`
		TxHash         string `db:"tx_hash"`
	}

	// MsgSetReferrer - db model for 'stwart_referrals_set_referrer'
	MsgSetReferrer struct {
		ID       uint64 `db:"id"`
		Creator  string `db:"creator"`
		Referrer string `db:"referrer"`
		Referral string `db:"referrals"`
		TxHash   string `db:"tx_hash"`
	}
)
