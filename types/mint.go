/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

import minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"

// MintParams represents the x/mint parameters
type MintParams struct {
	minttypes.Params
	Height int64
}

// NewMintParams allows to build a new MintParams instance
func NewMintParams(params minttypes.Params, height int64) *MintParams {
	return &MintParams{
		Params: params,
		Height: height,
	}
}
