/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountBalance represents the balance of an account at a given height
type AccountBalance struct {
	Address string
	Balance sdk.Coins
	Height  int64
}

// NewAccountBalance allows to build a new AccountBalance instance
func NewAccountBalance(address string, balance sdk.Coins, height int64) AccountBalance {
	return AccountBalance{
		Address: address,
		Balance: balance,
		Height:  height,
	}
}

// NativeTokenAmount represents the native token balance of an account at a given height
type NativeTokenAmount struct {
	Address string
	Balance math.Int
	Height  int64
}

// NewNativeTokenAmount allows to build a new NativeTokenAmount instance
func NewNativeTokenAmount(address string, balance math.Int, height int64) NativeTokenAmount {
	return NativeTokenAmount{
		Address: address,
		Balance: balance,
		Height:  height,
	}
}
