/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	"database/sql"
	"errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// GetAllMsgUpdateAddresses - method that get data from a db (stwart_secured_update_addresses).
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

// InsertMsgUpdateAddresses - insert a new MsgUpdateAddresses in a database (stwart_secured_update_addresses).
func (r Repository) InsertMsgUpdateAddresses(hash string, msgs ...*types.MsgUpdateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_secured_update_addresses (
			 id, tx_hash, creator, addresses
		) VALUES (
			$1, $2, $3, $4
		) RETURNING
			id, tx_hash, creator, addresses
	`

	for _, msg := range msgs {
		m, err := toMsgUpdateAddressesDatabase(hash, msg)
		if err != nil {
			return err
		}

		if _, err = r.db.Exec(q, m.ID, m.TxHash, m.Creator, m.Addresses); err != nil {
			if chain.IsAlreadyExists(err) {
				continue
			}
			return errs.Internal{Cause: err.Error()}
		}
	}

	return nil
}
