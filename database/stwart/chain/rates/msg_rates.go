/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// GetAllMsgCreateRates - method that get data from a db (stwartchain_rates_create_rates).
func (r Repository) GetAllMsgCreateRates(filter filter.Filter) ([]types.MsgCreateRates, error) {
	query, args := filter.Build(tableCreateRates)

	var result []MsgCreateRates
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableCreateRates}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableCreateRates}
	}

	return toMsgCreateRatesDomainList(result), nil
}

// InsertMsgCreateRates - insert a new MsgCreateRates in a database (stwartchain_rates_create_rates).
func (r Repository) InsertMsgCreateRates(hash string, msgs ...*types.MsgCreateRates) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_create_rates (
			creator, decimals, denom, rate, tx_hash
		) VALUES (
			$1, $2, $3, $4, $5
		) RETURNING
			id, creator, decimals, denom, rate, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgCreateRatesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.Decimals, m.Denom, m.Rate, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

func (r Repository) GetAllMsgUpdateRates(filter filter.Filter) ([]types.MsgUpdateRates, error) {
	query, args := filter.Build(tableUpdateRates)

	var result []MsgUpdateRates
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableUpdateRates}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableUpdateRates}
	}

	return toMsgUpdateRatesDomainList(result), nil
}

func (r Repository) InsertMsgUpdateRates(hash string, msgs ...*types.MsgUpdateRates) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_update_rates (
			creator, denom, rate, tx_hash
		) VALUES (
			$1, $2, $3, $4
		) RETURNING
			id, creator, denom, rate, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgUpdateRatesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.Denom, m.Rate, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

func (r Repository) GetAllMsgDeleteRates(filter filter.Filter) ([]types.MsgDeleteRates, error) {
	query, args := filter.Build(tableDeleteRates)

	var result []MsgDeleteRates
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableDeleteRates}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableDeleteRates}
	}

	return toMsgDeleteRatesDomainList(result), nil
}

func (r Repository) InsertMsgDeleteRates(hash string, msgs ...*types.MsgDeleteRates) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_delete_rates (
			creator, denom, tx_hash
		) VALUES (
			$1, $2, $3
		) RETURNING
			id, creator, denom, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgDeleteRatesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.Denom, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
