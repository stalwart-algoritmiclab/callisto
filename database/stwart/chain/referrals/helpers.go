/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/referrals"

// toMsgCreateUserDatabase converts MsgCreateUser to MsgCreateUserDatabase
func toMsgCreateUserDatabase(hash string, msg *referrals.MsgCreateUser) MsgCreateUser {
	return MsgCreateUser{
		Creator:        msg.Creator,
		AccountAddress: msg.AccountAddress,
		Referrer:       msg.Referrer,
		Referrals:      msg.Referrals,
		TxHash:         hash,
	}
}

// toMsgUpdateUserDatabase converts MsgUpdateUser to MsgUpdateUserDatabase
func toMsgUpdateUserDatabase(hash string, msg *referrals.MsgUpdateUser) MsgUpdateUser {
	return MsgUpdateUser{
		Creator:        msg.Creator,
		AccountAddress: msg.AccountAddress,
		Referrer:       msg.Referrer,
		Referrals:      msg.Referrals,
		TxHash:         hash,
	}
}

// toMsgDeleteUserDatabase converts MsgDeleteUser to MsgDeleteUserDatabase
func toMsgDeleteUserDatabase(hash string, msg *referrals.MsgDeleteUser) MsgDeleteUser {
	return MsgDeleteUser{
		Creator:        msg.Creator,
		AccountAddress: msg.AccountAddress,
		TxHash:         hash,
	}
}

// toMsgSetReferrerDatabase converts MsgSetReferrer to MsgSetReferrerDatabase
func toMsgSetReferrerDatabase(hash string, msg *referrals.MsgSetReferrer) MsgSetReferrer {
	return MsgSetReferrer{
		Creator:  msg.Creator,
		Referrer: msg.ReferrerAddress,
		Referral: msg.ReferrerAddress,
		TxHash:   hash,
	}
}