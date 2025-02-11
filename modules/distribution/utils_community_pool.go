/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package distribution

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

// updateCommunityPool fetch total amount of coins in the system from RPC and store it into database
func (m *Module) updateCommunityPool(height int64) error {
	log.Debug().Str("module", "distribution").Int64("height", height).Msg("getting community pool")

	pool, err := m.source.CommunityPool(height)
	if err != nil {
		return fmt.Errorf("error while getting community pool: %s", err)
	}

	// Store the pool into the database
	return m.db.SaveCommunityPool(pool, height)
}
