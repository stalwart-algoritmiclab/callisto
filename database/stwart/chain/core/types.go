/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

const (
	tableIssue    = "stwart_core_issue"
	tableWithdraw = "stwart_core_withdraw"
	tableSend     = "stwart_core_send"
)

type (
	// MsgIssue - db model for 'stwart_core_issue'
	MsgIssue struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Denom   string `db:"denom"`
		Amount  string `db:"amount"`
		Address string `db:"address"`
	}

	// MsgSend - db model for 'stwart_core_send'
	MsgSend struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		From    string `db:"from_address"`
		To      string `db:"to_address"`
		Amount  string `db:"amount"`
		Denom   string `db:"denom"`
	}

	// MsgWithdraw - db model for 'stwart_core_withdraw'
	MsgWithdraw struct {
		ID      uint64 `db:"id"`
		TxHash  string `db:"tx_hash"`
		Creator string `db:"creator"`
		Amount  string `db:"amount"`
		Denom   string `db:"denom"`
		Address string `db:"address"`
	}
)
