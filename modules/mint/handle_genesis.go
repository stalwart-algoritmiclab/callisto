/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package mint

import (
	"encoding/json"
	"fmt"

	tmtypes "github.com/cometbft/cometbft/types"

	"github.com/stalwart-algoritmiclab/callisto/types"

	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/rs/zerolog/log"
)

// HandleGenesis implements modules.Module
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", "mint").Msg("parsing genesis")

	// Read the genesis state
	var genState minttypes.GenesisState
	err := m.cdc.UnmarshalJSON(appState[minttypes.ModuleName], &genState)
	if err != nil {
		return fmt.Errorf("error while reading mint genesis data: %s", err)
	}

	// Save the params
	err = m.db.SaveMintParams(types.NewMintParams(genState.Params, doc.InitialHeight))
	if err != nil {
		return fmt.Errorf("error while storing genesis mint params: %s", err)
	}

	return nil
}
