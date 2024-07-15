/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
)

// InsertMsgCreateAddresses - insert a new MsgCreateAddresses in a database (stwartchain_feepolicy_create_addresses).
func (r Repository) InsertMsgCreateAddresses(hash string, msgs ...*types.MsgCreateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_feepolicy_create_addresses (
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

// InsertMsgUpdateAddresses - insert a new MsgUpdateAddresses in a database (stwartchain_feepolicy_update_addresses).
func (r Repository) InsertMsgUpdateAddresses(hash string, msgs ...*types.MsgUpdateAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_feepolicy_update_addresses (
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

// InsertMsgDeleteAddresses - insert a new MsgDeleteAddresses in a database (stwartchain_feepolicy_delete_addresses).
func (r Repository) InsertMsgDeleteAddresses(hash string, msgs ...*types.MsgDeleteAddresses) error {
	if len(msgs) == 0 || hash == "" {
		return nil
	}

	q := `
		INSERT INTO stwart_feepolicy_delete_addresses (
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
