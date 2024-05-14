/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package local

import (
	"fmt"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/forbole/juno/v6/node/local"

	mintsource "github.com/forbole/callisto/v4/modules/mint/source"
)

var (
	_ mintsource.Source = &Source{}
)

// Source implements mintsource.Source using a local node
type Source struct {
	*local.Source
	querier minttypes.QueryServer
}

// NewSource returns a new Source instance
func NewSource(source *local.Source, querier minttypes.QueryServer) *Source {
	return &Source{
		Source:  source,
		querier: querier,
	}
}

// GetInflation implements mintsource.Source
func (s Source) GetInflation(height int64) (math.LegacyDec, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return math.LegacyDec{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.Inflation(sdk.WrapSDKContext(ctx), &minttypes.QueryInflationRequest{})
	if err != nil {
		return math.LegacyDec{}, err
	}

	return res.Inflation, nil
}

// Params implements mintsource.Source
func (s Source) Params(height int64) (minttypes.Params, error) {
	ctx, err := s.LoadHeight(height)
	if err != nil {
		return minttypes.Params{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.querier.Params(sdk.WrapSDKContext(ctx), &minttypes.QueryParamsRequest{})
	if err != nil {
		return minttypes.Params{}, err
	}

	return res.Params, nil
}
