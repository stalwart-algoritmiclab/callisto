/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch feepolicyMsg := msg.(type) {
	case *feepolicy.MsgCreateTariffs:
		return m.handleMsgCreateTariffs(tx, index, feepolicyMsg)
	case *feepolicy.MsgUpdateTariffs:
		return m.handleMsgUpdateTariffs(tx, index, feepolicyMsg)
	case *feepolicy.MsgDeleteTariffs:
		return m.handleMsgDeleteTariffs(tx, index, feepolicyMsg)
	case *feepolicy.MsgCreateAddresses:
		return m.handleMsgCreateAddresses(tx, index, feepolicyMsg)
	case *feepolicy.MsgUpdateAddresses:
		return m.handleMsgUpdateAddresses(tx, index, feepolicyMsg)
	case *feepolicy.MsgDeleteAddresses:
		return m.handleMsgDeleteAddresses(tx, index, feepolicyMsg)
	default:
		return nil
	}
}
