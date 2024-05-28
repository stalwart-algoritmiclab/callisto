/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package exchanger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg sdk.Msg, tx *juno.Tx) error {
	//if len(tx.Logs) == 0 {
	//	return nil
	//}

	switch exchangerMsg := msg.(type) {
	case *exchanger.MsgExchange:
		return m.handleMsgExchange(tx, index, exchangerMsg)
	default:
		return nil
	}
}
