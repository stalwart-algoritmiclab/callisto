/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package handlers

import (
	"fmt"

	"github.com/forbole/callisto/modules/actions/types"

	"github.com/rs/zerolog/log"
)

func ValidatorCommissionAmountHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("address", payload.GetAddress()).
		Int64("height", payload.Input.Height).
		Msg("executing validator commission action")

	// Get latest node height
	height, err := ctx.GetHeight(nil)
	if err != nil {
		return nil, err
	}

	// Get validator total commission value
	commission, err := ctx.Sources.DistrSource.ValidatorCommission(payload.GetAddress(), height)
	if err != nil {
		return nil, fmt.Errorf("error while getting validator commission: %s", err)
	}

	return types.ValidatorCommissionAmount{
		Coins: types.ConvertDecCoins(commission),
	}, nil
}
