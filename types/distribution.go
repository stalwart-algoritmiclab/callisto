/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

import (
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// DistributionParams represents the parameters of the x/distribution module
type DistributionParams struct {
	distrtypes.Params
	Height int64
}

// NewDistributionParams allows to build a new DistributionParams instance
func NewDistributionParams(params distrtypes.Params, height int64) *DistributionParams {
	return &DistributionParams{
		Params: params,
		Height: height,
	}
}
