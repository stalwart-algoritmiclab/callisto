/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package distribution

import (
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/forbole/callisto/v4/types"
)

// UpdateParams gets the updated params and stores them inside the database
func (m *Module) UpdateParams(height int64) error {
	log.Debug().Str("module", "distribution").Int64("height", height).
		Msg("updating params")

	params, err := m.source.Params(height)
	if err != nil {
		return fmt.Errorf("error while getting params: %s", err)
	}

	return m.db.SaveDistributionParams(types.NewDistributionParams(params, height))

}
