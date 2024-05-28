package core

import (
	"encoding/json"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
)

const moduleName = "core"

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", moduleName).Msg("parsing genesis")

	// Unmarshal the faucet state
	var exchangerState exchanger.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[moduleName], &exchangerState); err != nil {
		return err
	}

	return nil
}
