package rates

import (
	"encoding/json"

	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
)

const moduleName = "rates"

// HandleGenesis implements GenesisModule
func (m *Module) HandleGenesis(doc *tmtypes.GenesisDoc, appState map[string]json.RawMessage) error {
	log.Debug().Str("module", moduleName).Msg("parsing genesis")

	// Unmarshal the rates state
	var ratesState rates.GenesisState
	if err := m.cdc.UnmarshalJSON(appState[moduleName], &ratesState); err != nil {
		return err
	}

	return nil
}
