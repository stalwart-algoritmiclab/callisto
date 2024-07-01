/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package secured

import (
	"github.com/lib/pq"
)

const (
	tableCreateAddresses = "stwart_secured_create_addresses"
	tableDeleteAddresses = "stwart_secured_delete_addresses"
	tableUpdateAddresses = "stwart_secured_update_addresses"
)

// MsgCreateAddresses - db model for 'stwart_secured_create_addresses'
type (
	MsgCreateAddresses struct {
		ID        uint64         `db:"id"`
		Creator   string         `db:"creator"`
		TxHash    string         `db:"tx_hash"`
		Addresses pq.StringArray `db:"addresses"`
	}

	// MsgDeleteAddresses - db model for 'stwart_secured_delete_addresses'
	MsgDeleteAddresses struct {
		ID        uint64 `db:"id"`         // Auto increment ID
		AddressID uint64 `db:"address_id"` // ID in message
		Creator   string `db:"creator"`
		TxHash    string `db:"tx_hash"`
	}

	// MsgUpdateAddresses - db model for 'stwart_secured_update_addresses'
	MsgUpdateAddresses struct {
		ID        uint64         `db:"id"`         // Auto increment ID
		AddressID uint64         `db:"address_id"` // ID in message
		Creator   string         `db:"creator"`
		TxHash    string         `db:"tx_hash"`
		Addresses pq.StringArray `db:"addresses"`
	}
)
