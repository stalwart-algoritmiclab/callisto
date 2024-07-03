/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	"database/sql"
	"errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// GetAllMsgCreateAddresses - method that get data from a db (stwart_secured_create_addresses).
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

// InsertMsgCreateAddresses - insert a new MsgCreateAddresses in a database (stwart_secured_create_addresses).
func (r Repository) InsertMsgCreateAddresses(hash string, msgs ...*types.MsgCreateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_secured_create_addresses (
			 tx_hash, creator, addresses
		) VALUES (
			$1, $2, $3
		) RETURNING
			id, tx_hash, creator, addresses
	`

	for _, msg := range msgs {
		m, err := toMsgCreateAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.TxHash, m.Creator, m.Addresses); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
