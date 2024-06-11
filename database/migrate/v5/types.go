/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package v5

type MessageRow struct {
	TransactionHash           string `db:"transaction_hash"`
	Index                     int64  `db:"index"`
	Type                      string `db:"type"`
	Value                     string `db:"value"`
	InvolvedAccountsAddresses string `db:"involved_accounts_addresses"`
	Height                    int64  `db:"height"`
}
