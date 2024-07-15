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

// GetAllMsgCreateAddresses - method that get data from a db (stwartchain_rates_create_addresses).
func (r Repository) GetAllMsgCreateAddresses(filter filter.Filter) ([]types.MsgCreateAddresses, error) {
	query, args := filter.Build(tableCreateAddresses)

	var result []MsgCreateAddresses
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableCreateAddresses}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableCreateAddresses}
	}

	return toMsgCreateAddressesDomainList(result), nil
}

// InsertMsgCreateAddresses - insert a new MsgCreateAddresses in a database (stwartchain_rates_create_addresses).
func (r Repository) InsertMsgCreateAddresses(hash string, msgs ...*types.MsgCreateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_create_addresses (
			 creator, address, tx_hash
		) VALUES (
			$1, $2, $3
		) RETURNING
			id,  creator, address, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgCreateAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.Address, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}

// GetAllMsgUpdateAddresses - method that get data from a db (stwartchain_rates_update_addresses).
func (r Repository) GetAllMsgUpdateAddresses(filter filter.Filter) ([]types.MsgUpdateAddresses, error) {
	query, args := filter.Build(tableUpdateAddresses)

	var result []MsgUpdateAddresses
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableUpdateAddresses}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableUpdateAddresses}
	}

	return toMsgUpdateAddressesDomainList(result), nil

}

// InsertMsgUpdateAddresses - insert a new MsgUpdateAddresses in a database (stwartchain_rates_update_addresses).
func (r Repository) InsertMsgUpdateAddresses(hash string, msgs ...*types.MsgUpdateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_update_addresses (
			creator, address, address_id, tx_hash
		) VALUES (
			$1, $2, $3, $4
		) RETURNING
			id, creator, address, address_id, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgUpdateAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.Address, m.AddressID, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil

}

// GetAllMsgDeleteAddresses - method that get data from a db (stwartchain_rates_delete_addresses).
func (r Repository) GetAllMsgDeleteAddresses(filter filter.Filter) ([]types.MsgDeleteAddresses, error) {
	query, args := filter.Build(tableDeleteAddresses)

	var result []MsgDeleteAddresses
	if err := r.db.Select(&result, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NotFound{What: tableDeleteAddresses}
		}

		return nil, errs.Internal{Cause: err.Error()}
	}
	if len(result) == 0 {
		return nil, errs.NotFound{What: tableDeleteAddresses}
	}

	return toMsgDeleteAddressesDomainList(result), nil
}

// InsertMsgDeleteAddresses - insert a new MsgDeleteAddresses in a database (stwartchain_rates_delete_addresses).
func (r Repository) InsertMsgDeleteAddresses(hash string, msgs ...*types.MsgDeleteAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_rates_delete_addresses (
			creator, address_id, tx_hash
		) VALUES (
			$1, $2, $3
		) RETURNING
			id, creator, address_id, tx_hash
	`

	for _, msg := range msgs {
		m, err := toMsgDeleteAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err := r.db.Exec(q, m.Creator, m.AddressID, m.TxHash); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
