/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
)

// handleMsgUpdateTariffs allows to properly handle a MsgUpdateTariffs
func (m *Module) handleMsgUpdateTariffs(tx *juno.Tx, _ int, msg *feepolicy.MsgUpdateTariffs) error {
	return m.feepolicyRepo.InsertMsgUpdateTariffs(tx.Height, tx.TxHash, msg)
}
