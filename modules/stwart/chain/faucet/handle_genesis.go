package faucet

import (
	"encoding/json"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
)

const moduleName = "faucet"

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", moduleName).Msg("parsing genesis")

	// Unmarshal the faucet state
	var faucetState faucet.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[moduleName], &faucetState); err != nil {
		return err
	}

	return nil // TODO: add table stats and methods
}
