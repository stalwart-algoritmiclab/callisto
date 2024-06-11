/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package exchanger

import (
	juno "github.com/forbole/juno/v5/types"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

// handleMsgExchange handles the MsgExchange message
func (m *Module) handleMsgExchange(tx *juno.Tx, _ int, msg *exchanger.MsgExchange) error {
	return m.exchangerRepo.InsertMsgExchange(tx.TxHash, msg)
}
