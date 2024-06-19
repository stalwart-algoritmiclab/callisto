/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
)

// InsertDeleteMsgDeleteTariffs - insert MsgDeleteTariffs in the database
func (r Repository) InsertDeleteMsgDeleteTariffs(height int64, hash string, msgs ...*feepolicy.MsgDeleteTariffs) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	for _, msg := range msgs {
		m := toMsgDeleteTariffsDataBase(msg)

		query := `INSERT INTO stwart_feepolicy_msg_delete_tariffs (denom, creator, tx_hash, tariff_id, fee_id, height) VALUES ($1, $2, $3, $4, $5, $6)`
		if _, err := r.db.Exec(query, m.Denom, m.Creator, hash, m.TariffID, m.FeeID, height); err != nil {
			if chain.IsAlreadyExists(err) {
				return nil
			}

			return err
		}
	}

	return nil
}
