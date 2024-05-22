/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package slashing

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/types"
)

// UpdateParams gets the slashing params for the given height, and stores them inside the database
func (m *Module) UpdateParams(height int64) error {
	log.Debug().Str("module", "slashing").Int64("height", height).Msg("updating params")

	params, err := m.source.GetParams(height)
	if err != nil {
		return fmt.Errorf("error while getting params: %s", err)
	}

	return m.db.SaveSlashingParams(types.NewSlashingParams(params, height))

}
