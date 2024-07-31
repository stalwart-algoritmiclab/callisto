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

// handleMsgCreatePollsParams allows to properly handle a MsgCreatePollsParams
func (m *Module) handleMsgCreatePollsParams(tx *juno.Transaction, _ int, msg *types.MsgCreatePollsParams) error {
	return m.pollsRepo.InsertMsgCreatePollsParams(tx.TxHash, msg)
}

// handleMsgUpdatePollsParams allows to properly handle a MsgUpdatePollsParams
func (m *Module) handleMsgUpdatePollsParams(tx *juno.Transaction, _ int, msg *types.MsgUpdatePollsParams) error {
	return m.pollsRepo.InsertMsgUpdatePollsParams(tx.TxHash, msg)
}

// handleMsgDeletePollsParams allows to properly handle a MsgDeletePollsParams
func (m *Module) handleMsgDeletePollsParams(tx *juno.Transaction, _ int, msg *types.MsgDeletePollsParams) error {
	return m.pollsRepo.InsertMsgDeletePollsParams(tx.TxHash, msg)
}
