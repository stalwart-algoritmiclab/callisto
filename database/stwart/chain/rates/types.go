/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import "github.com/lib/pq"

const (
	tableCreateAdrresses = "stwart_rates_create_addresses"
	tableCreateRates     = "stwart_rates_create_rates"
)

// MsgCreateAddresses - db model for 'stwart_rates_create_addresses'
type MsgCreateAddresses struct {
	ID      uint64         `db:"id"`
	Creator string         `db:"creator"`
	Address pq.StringArray `db:"address"`
	TxHash  string         `db:"tx_hash"`
}

// MsgCreateRates - db model for 'stwart_rates_create_rates'
type MsgCreateRates struct {
	ID       uint64 `db:"id"`
	Creator  string `db:"creator"`
	Decimals int32  `db:"decimals"`
	Denom    string `db:"denom"`
	Rate     string `db:"rate"`
	TxHash   string `db:"tx_hash"`
}
