/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import sdk "github.com/cosmos/cosmos-sdk/types"

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg *MsgSetReferrer) GetSigners() []sdk.AccAddress {
	return msg.GetSigners()
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg *MsgCreateUser) GetSigners() []sdk.AccAddress {
	return msg.GetSigners()
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg *MsgUpdateUser) GetSigners() []sdk.AccAddress {
	return msg.GetSigners()
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg *MsgDeleteUser) GetSigners() []sdk.AccAddress {
	return msg.GetSigners()
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}
