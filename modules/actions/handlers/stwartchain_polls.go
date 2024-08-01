/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package handlers

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/modules/actions/types"
)

// StalwartChainPollsHandler is a handler function for the stalwart chain polls action
func StalwartChainPollsHandler(ctx *types.Context, payload *types.Payload) (interface{}, error) {
	log.Debug().Str("poll_id", strconv.FormatUint(payload.Input.PollID, 10)).
		Int64("height", payload.Input.Height).
		Msg("executing stalwart chain polls action")

	height, err := ctx.GetHeight(payload)
	if err != nil {
		return nil, err
	}

	// If the limit and offset are not set, we can get the one poll directly
	if payload.Input.Limit <= 0 && payload.Input.Offset <= 0 {
		poll, err := ctx.Sources.PollSource.GetPoll(
			payload.Input.PollID,
			height,
		)
		if err != nil {
			return nil, fmt.Errorf("error while getting poll: %s", err)
		}

		return types.PollsResponse{
			Polls: []types.Poll{types.ToPollResponse(poll.Polls)},
		}, nil
	}

	// Get all polls
	res, err := ctx.Sources.PollSource.GetAllPoll(height, payload.GetPagination())
	if err != nil {
		return nil, fmt.Errorf("error while getting polls: %s", err)
	}

	return types.PollsResponse{
		Polls:      types.ToPollsResponse(res.Polls),
		Pagination: res.Pagination,
	}, nil
}
