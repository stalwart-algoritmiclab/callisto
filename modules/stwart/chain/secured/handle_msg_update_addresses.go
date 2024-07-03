/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	juno "github.com/forbole/juno/v6/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"
)

// handleMsgIssue allows to properly handle a MsgIssue
func (m *Module) handleMsgUpdateAddresses(tx *juno.Transaction, _ int, msg *types.MsgUpdateAddresses) error {
	return m.securedRepo.InsertMsgUpdateAddresses(tx.TxHash, msg)
}
