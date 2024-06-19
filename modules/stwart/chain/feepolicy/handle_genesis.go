/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feepolicy

import (
	"encoding/json"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
)

const moduleName = "feepolicy"

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", moduleName).Msg("parsing genesis")

	// Unmarshal the feepolicy state
	var feepolicyState feepolicy.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[moduleName], &feepolicyState); err != nil {
		return err
	}

	return nil // TODO: add table stats and methods
}
