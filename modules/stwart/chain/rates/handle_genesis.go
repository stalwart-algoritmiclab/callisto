/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import (
	"encoding/json"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"
)

const moduleName = "rates"

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", moduleName).Msg("parsing genesis")

	// Unmarshal the rates state
	var ratesState types.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[moduleName], &ratesState); err != nil {
		return err
	}

	return nil
}
