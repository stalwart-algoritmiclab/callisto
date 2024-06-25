/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	juno "github.com/forbole/juno/v6/types"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
)

// handleMsgIssue allows to properly handle a MsgIssue
func (m *Module) handleMsgCreateAddresses(tx *juno.Transaction, _ int, msg *types.MsgCreateAddresses) error {
	return m.securedRepo.InsertMsgCreateAddresses(tx.TxHash, msg)
}
