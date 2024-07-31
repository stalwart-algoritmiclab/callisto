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

// handleMsgVote allows to properly handle a MsgVote
func (m *Module) handleMsgVote(tx *juno.Transaction, _ int, msg *types.MsgVote) error {
	return m.pollsRepo.InsertMsgVote(tx.TxHash, msg)
}
