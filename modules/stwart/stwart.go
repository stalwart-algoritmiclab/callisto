package stwart

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v5/logging"
	jmodules "github.com/forbole/juno/v5/modules"
	"github.com/forbole/juno/v5/node"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain/last_block"
	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger"
	exchangersource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source"
	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/faucet"
	faucetsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/faucet/source"

	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/core"
	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/rates"
	ratessource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/rates/source"
	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/secured"
	securedsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/secured/source"
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

	faucetSource faucetsource.Source,
	exchangerSource exchangersource.Source,
	securedSource securedsource.Source,
	ratesSource ratessource.Source,
	coreSource core.Source,
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
			faucet.NewModule(faucetSource, cdc, db),
			secured.NewModule(securedSource, cdc, db),
			rates.NewModule(ratesSource, cdc, db),
			core.NewModule(coreSource, cdc, db),
		},
	}

	go m.scheduler()

	return m
}

// Name implements modules.Module
func (m *Module) Name() string { return module }
