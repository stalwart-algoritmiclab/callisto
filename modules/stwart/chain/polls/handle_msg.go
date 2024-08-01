/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package polls

import (
	juno "github.com/forbole/juno/v6/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"

	"github.com/stalwart-algoritmiclab/callisto/utils"
)

// msgFilter defines the messages that should be handled by this module
var msgFilter = map[string]bool{
	"/stwartchain.polls.MsgCreatePollsParams": true,
	"/stwartchain.polls.MsgUpdatePollsParams": true,
	"/stwartchain.polls.MsgDeletePollsParams": true,
	"/stwartchain.polls.MsgCreatePoll":        true,
	"/stwartchain.polls.MsgVote":              true,
}

// HandleMsg implements MessageModule
func (m *Module) HandleMsg(index int, msg juno.Message, tx *juno.Transaction) error {
	if _, ok := msgFilter[msg.GetType()]; !ok {
		return nil
	}

	switch msg.GetType() {
	case "/stwartchain.polls.MsgCreatePollsParams":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreatePollsParams{})
		return m.handleMsgCreatePollsParams(tx, index, cosmosMsg)
	case "/stwartchain.polls.MsgUpdatePollsParams":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgUpdatePollsParams{})
		return m.handleMsgUpdatePollsParams(tx, index, cosmosMsg)
	case "/stwartchain.polls.MsgDeletePollsParams":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgDeletePollsParams{})
		return m.handleMsgDeletePollsParams(tx, index, cosmosMsg)
	case "/stwartchain.polls.MsgCreatePoll":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgCreatePoll{})
		return m.handleMsgCreatePoll(tx, index, cosmosMsg)
	case "/stwartchain.polls.MsgVote":
		cosmosMsg := utils.UnpackMessage(m.cdc, msg.GetBytes(), &types.MsgVote{})
		return m.handleMsgVote(tx, index, cosmosMsg)
	default:
		return nil
	}
}
