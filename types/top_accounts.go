/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

const (
	TokenDenomSTW        = "stw"
	DecimalTokenDenomSTW = 1_0000_0000
)

// TopAccount represents a stwart account from top accounts module
type TopAccount struct {
	Address string
	Type    string
}

// NewTopAccount allows to build a new TopAccount instance
func NewTopAccount(address, accountType string) TopAccount {
	return TopAccount{
		Address: address,
		Type:    accountType,
	}
}
