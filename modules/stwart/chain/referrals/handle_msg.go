/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package referrals

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/referrals"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(_ int, msg sdk.Msg, tx *juno.Tx) error {
	switch referralsMsg := msg.(type) {
	case *referrals.MsgCreateUser:
		return m.handleMsgCreateUser(tx, referralsMsg)
	case *referrals.MsgUpdateUser:
		return m.handleMsgUpdateUser(tx, referralsMsg)
	case *referrals.MsgDeleteUser:
		return m.handleMsgDeleteUser(tx, referralsMsg)
	case *referrals.MsgSetReferrer:
		return m.handleMsgSetReferrer(tx, referralsMsg)

	default:
		return nil
	}
}
