package modules

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/stalwart-algoritmiclab/callisto/modules/actions"
	"github.com/stalwart-algoritmiclab/callisto/modules/types"

	"github.com/forbole/juno/v6/modules/pruning"
	"github.com/forbole/juno/v6/modules/telemetry"

	"github.com/forbole/callisto/v4/modules/slashing"

	jmodules "github.com/forbole/juno/v6/modules"
	"github.com/forbole/juno/v6/modules/messages"
	"github.com/forbole/juno/v6/modules/registrar"

	"github.com/stalwart-algoritmiclab/callisto/utils"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/modules/auth"
	"github.com/stalwart-algoritmiclab/callisto/modules/bank"
	"github.com/stalwart-algoritmiclab/callisto/modules/consensus"
	"github.com/stalwart-algoritmiclab/callisto/modules/distribution"
	"github.com/stalwart-algoritmiclab/callisto/modules/feegrant"

	juno "github.com/forbole/juno/v6/types"

	dailyrefetch "github.com/stalwart-algoritmiclab/callisto/modules/daily_refetch"
	"github.com/stalwart-algoritmiclab/callisto/modules/gov"
	messagetype "github.com/stalwart-algoritmiclab/callisto/modules/message_type"
	"github.com/stalwart-algoritmiclab/callisto/modules/mint"
	"github.com/stalwart-algoritmiclab/callisto/modules/modules"
	"github.com/stalwart-algoritmiclab/callisto/modules/pricefeed"
	"github.com/stalwart-algoritmiclab/callisto/modules/staking"
	"github.com/stalwart-algoritmiclab/callisto/modules/upgrade"
)

// UniqueAddressesParser returns a wrapper around the given parser that removes all duplicated addresses
func UniqueAddressesParser(parser messages.MessageAddressesParser) messages.MessageAddressesParser {
	return func(tx *juno.Transaction) ([]string, error) {
		addresses, err := parser(tx)
		if err != nil {
			return nil, err
		}

		return utils.RemoveDuplicateValues(addresses), nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

var (
	_ registrar.Registrar = &Registrar{}
)

// Registrar represents the modules.Registrar that allows to register all modules that are supported by BigDipper
type Registrar struct {
	parser messages.MessageAddressesParser
	cdc    codec.Codec
}

// NewRegistrar allows to build a new Registrar instance
func NewRegistrar(parser messages.MessageAddressesParser, cdc codec.Codec) *Registrar {
	return &Registrar{
		parser: UniqueAddressesParser(parser),
		cdc:    cdc,
	}
}

// BuildModules implements modules.Registrar
func (r *Registrar) BuildModules(ctx registrar.Context) jmodules.Modules {
	db := database.Cast(ctx.Database)

	sources, err := types.BuildSources(ctx.JunoConfig.Node, r.cdc)
	if err != nil {
		panic(err)
	}

	actionsModule := actions.NewModule(ctx.JunoConfig, r.cdc, sources)
	authModule := auth.NewModule(r.parser, r.cdc, db)
	bankModule := bank.NewModule(r.parser, sources.BankSource, r.cdc, db)
	consensusModule := consensus.NewModule(db)
	dailyRefetchModule := dailyrefetch.NewModule(ctx.Proxy, db)
	distrModule := distribution.NewModule(sources.DistrSource, r.cdc, db)
	feegrantModule := feegrant.NewModule(r.cdc, db)
	messagetypeModule := messagetype.NewModule(r.parser, r.cdc, db)
	mintModule := mint.NewModule(sources.MintSource, r.cdc, db)
	slashingModule := slashing.NewModule(sources.SlashingSource, r.cdc, db)
	stakingModule := staking.NewModule(sources.StakingSource, r.cdc, db)
	govModule := gov.NewModule(sources.GovSource, distrModule, mintModule, slashingModule, stakingModule, r.cdc, db)
	upgradeModule := upgrade.NewModule(db, stakingModule)
	stwartModule := stwart.NewModule(
		cdc,
		db,
		ctx.Proxy,
		ctx.Logger,
		sources.FaucetSource,
		sources.ExchangerSource,
	)

	return []jmodules.Module{
		messages.NewModule(r.parser, ctx.Database),
		telemetry.NewModule(ctx.JunoConfig),
		pruning.NewModule(ctx.JunoConfig, db, ctx.Logger),

		actionsModule,
		authModule,
		bankModule,
		consensusModule,
		dailyRefetchModule,
		distrModule,
		feegrantModule,
		govModule,
		mintModule,
		messagetypeModule,
		modules.NewModule(ctx.JunoConfig.Chain, db),
		pricefeed.NewModule(ctx.JunoConfig, r.cdc, db),
		slashingModule,
		stakingModule,
		upgradeModule,
		stwartModule,
	}
}
