/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package faucet

const (
	tableFaucet = "stwart_faucet"
)

// MsgIssue - db model for 'stwart_faucet'
type MsgIssue struct {
	ID      uint64 `db:"id"`
	Creator string `db:"creator"`
	Address string `db:"address"`
	TxHash  string `db:"tx_hash"`
}
