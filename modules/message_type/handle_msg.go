/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package message_type

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"

	"github.com/stalwart-algoritmiclab/callisto/modules/utils"
	msgtypes "github.com/stalwart-algoritmiclab/callisto/types"

	"github.com/forbole/juno/v6/types"
)

// HandleMsg represents a message handler that stores the given message inside the proper database table
func (m *Module) HandleMsg(
	index int, msg types.Message, tx *types.Transaction) error {
	// Save message type
	err := m.db.SaveMessageType(msgtypes.NewMessageType(
		msg.GetType(),
		utils.GetModuleNameFromTypeURL(msg.GetType()),
		utils.GetMsgFromTypeURL(msg.GetType()),
		int64(tx.Height)))

	if err != nil {
		return err
	}

	return nil
}
