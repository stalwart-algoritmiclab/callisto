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

// GetAllMsgDeleteAddresses - method that get data from a db (stwart_secured_delete_addresses).
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

// InsertMsgDeleteAddresses - insert a new MsgDeleteAddresses in a database (stwart_secured_delete_addresses).
func (r Repository) InsertMsgDeleteAddresses(hash string, msgs ...*types.MsgDeleteAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_secured_delete_addresses (
		    address_id, tx_hash, creator
		) VALUES (
			$1, $2, $3
		) RETURNING
			id, address_id, tx_hash, creator
	`

	for _, msg := range msgs {
		m, err := toMsgDeleteAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.AddressID, m.TxHash, m.Creator); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
