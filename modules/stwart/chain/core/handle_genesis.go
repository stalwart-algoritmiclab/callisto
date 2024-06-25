/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import (
	"encoding/json"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

const moduleName = "core"

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", moduleName).Msg("parsing genesis")

	// Unmarshal the faucet state
	var coreState types.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[moduleName], &coreState); err != nil {
		return err
	}

	return nil
}
