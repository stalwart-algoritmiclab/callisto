package secured

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v5/modules"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain/secured"
	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/secured/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/secured module
type Module struct {
	cdc         codec.Codec
	db          *database.Db
	securedRepo secured.Repository

	keeper source.Source
}

// NewModule returns a new Module instance
func NewModule(keeper source.Source, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		keeper:      keeper,
		cdc:         cdc,
		db:          db,
		securedRepo: *secured.NewRepository(db.Sqlx, cdc),
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "stwart_secured"
}