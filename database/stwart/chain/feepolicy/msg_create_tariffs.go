/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
)

// InsertMsgCreateTariffs - insert MsgCreateTariffs into the database
func (r Repository) InsertMsgCreateTariffs(height int64, hash string, msgs ...*feepolicy.MsgCreateTariffs) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	for _, msg := range msgs {
		if err := r.addMsgCreateTariff(hash, toTariffDatabase(msg.Tariffs)); err != nil {
			return errs.Internal{Cause: err.Error()}
		}

		if err := r.addMsgCreateTariffs(height, hash, toMsgCreateTariffsDataBase(msg)); err != nil {
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// addMsgCreateTariffs - create tariffs in the database
func (r Repository) addMsgCreateTariffs(height int64, hash string, tariff MsgCreateTariffs) error {
	query := `INSERT INTO stwart_feepolicy_msg_create_tariffs (denom, creator, tx_hash, height) VALUES ($1, $2, $3, $4)`
	if _, err := r.db.Exec(query, tariff.Denom, tariff.Creator, hash, height); err != nil {
		if chain.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}

// addTariff - create tariff in the database
func (r Repository) addMsgCreateTariff(hash string, tariffDetail Tariff) error {
	query := `INSERT INTO stwart_feepolicy_tariffs (tariff_id, amount, denom, min_ref_balance, fees, tx_hash) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	if _, err := r.db.Exec(query, tariffDetail.TariffID, tariffDetail.Amount, tariffDetail.Denom, tariffDetail.MinRefBalance, tariffDetail.Fees, hash); err != nil {
		if chain.IsAlreadyExists(err) {
			return nil
		}

		return err
	}

	return nil
}
