/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package exchanger

const (
	tableExchange = "stwart_exchanger"
)

// MsgExchange - db model for 'stwart_exchanger'
type MsgExchange struct {
	ID      uint64 `db:"id"`
	Creator string `db:"creator"`
	Denom   string `db:"denom"`
	Amount  string `db:"amount"`
	DenomTo string `db:"denom_to"`
	TxHash  string `db:"tx_hash"`
}
