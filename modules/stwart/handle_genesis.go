/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package stwart

import (
	"encoding/json"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/forbole/juno/v5/modules"
)

// HandleGenesis implements modules.GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	for _, module := range m.stwartModules {
		if genesisModule, ok := module.(modules.GenesisModule); ok {
			if err := genesisModule.HandleGenesis(doc, appState); err != nil {
				m.logger.GenesisError(module, err)
			}
		}
	}

	return nil
}
