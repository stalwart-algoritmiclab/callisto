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

// handleMsgDeleteTariffs allows to properly handle a MsgDeleteTariffs
func (m *Module) handleMsgDeleteTariffs(tx *juno.Tx, _ int, msg *feepolicy.MsgDeleteTariffs) error {
	return m.feepolicyRepo.InsertDeleteMsgDeleteTariffs(tx.Height, tx.TxHash, msg)
}
