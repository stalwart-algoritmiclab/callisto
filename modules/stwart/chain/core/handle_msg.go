/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/core"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	switch someMsg := msg.(type) {
	case *core.MsgWithdraw:
		return m.handleMsgWithdraw(tx, index, someMsg)
	case *core.MsgIssue:
		return m.handleMsgIssue(tx, index, someMsg)
	case *core.MsgRefund:
		return m.handleMsgRefund(tx, index, someMsg)
	case *core.MsgRefReward:
		return m.handleMsgRefReward(tx, index, someMsg)
	case *core.MsgSend:
		return m.handleMsgSend(tx, index, someMsg)
	case *core.MsgFees:
		return m.handleMsgFees(tx, index, someMsg)

	default:
		return nil
	}
}
