/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import (
	juno "github.com/forbole/juno/v6/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

// handleMsgCreatePoll allows to properly handle a MsgCreatePoll
func (m *Module) handleMsgCreatePoll(tx *juno.Transaction, _ int, msg *types.MsgCreatePoll) error {
	return m.pollsRepo.InsertMsgCreatePoll(tx.TxHash, msg)
}
