/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch securedMsg := msg.(type) {
	case *secured.MsgCreateAddresses:
		return m.handleMsgCreateAddresses(tx, index, securedMsg)
	case *secured.MsgDeleteAddresses:
		return m.handleMsgDeleteAddresses(tx, index, securedMsg)
	case *secured.MsgUpdateAddresses:
		return m.handleMsgUpdateAddresses(tx, index, securedMsg)
	default:
		return nil
	}
}
