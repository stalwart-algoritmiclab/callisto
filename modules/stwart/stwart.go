package stwart

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v5/logging"
	jmodules "github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/node"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain/last_block"
	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger"
	exchangerSource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source"
)

var (
	_ jmodules.Module        = &Module{}
	_ jmodules.GenesisModule = &Module{}
)

type stwartModule interface {
	jmodules.Module
	jmodules.GenesisModule
	jmodules.MessageModule
}

type Module struct {
	cdc           codec.Codec
	db            *database.Db
	lastBlockRepo last_block.Repository
	logger        logging.Logger
	stwartModules []stwartModule
	node          node.Node
}

func NewModule(
	cdc codec.Codec,
	db *database.Db,
	node node.Node,
	logger logging.Logger,

	exchangerSource exchangerSource.Source,
) *Module {
	m := &Module{
		cdc:           cdc,
		db:            db,
		lastBlockRepo: *last_block.NewRepository(db.Sqlx),
		node:          node,
		logger:        logger,
		stwartModules: []stwartModule{
			// stwart modules
			exchanger.NewModule(exchangerSource, cdc, db),
		},
	}

	go m.scheduler()

	return m
}

// Name implements modules.Module
func (m *Module) Name() string { return module }
