/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

// handleMsgIssue allows to properly handle a MsgIssue
func (m *Module) handleMsgCreateAddresses(tx *juno.Tx, _ int, msg *secured.MsgCreateAddresses) error {
	return m.securedRepo.InsertMsgCreateAddresses(tx.TxHash, msg)
}
