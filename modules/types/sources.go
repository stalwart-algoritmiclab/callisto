package types

import (
	"fmt"
	"os"

	"cosmossdk.io/simapp"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/forbole/juno/v5/node/remote"
	"github.com/forbole/juno/v5/types/params"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/juno/v5/node/local"

	nodeconfig "github.com/forbole/juno/v5/node/config"

	banksource "github.com/stalwart-algoritmiclab/callisto/modules/bank/source"
	localbanksource "github.com/stalwart-algoritmiclab/callisto/modules/bank/source/local"
	remotebanksource "github.com/stalwart-algoritmiclab/callisto/modules/bank/source/remote"
	distrsource "github.com/stalwart-algoritmiclab/callisto/modules/distribution/source"
	remotedistrsource "github.com/stalwart-algoritmiclab/callisto/modules/distribution/source/remote"
	govsource "github.com/stalwart-algoritmiclab/callisto/modules/gov/source"
	localgovsource "github.com/stalwart-algoritmiclab/callisto/modules/gov/source/local"
	remotegovsource "github.com/stalwart-algoritmiclab/callisto/modules/gov/source/remote"
	mintsource "github.com/stalwart-algoritmiclab/callisto/modules/mint/source"
	localmintsource "github.com/stalwart-algoritmiclab/callisto/modules/mint/source/local"
	remotemintsource "github.com/stalwart-algoritmiclab/callisto/modules/mint/source/remote"
	slashingsource "github.com/stalwart-algoritmiclab/callisto/modules/slashing/source"
	localslashingsource "github.com/stalwart-algoritmiclab/callisto/modules/slashing/source/local"
	remoteslashingsource "github.com/stalwart-algoritmiclab/callisto/modules/slashing/source/remote"
	stakingsource "github.com/stalwart-algoritmiclab/callisto/modules/staking/source"
	localstakingsource "github.com/stalwart-algoritmiclab/callisto/modules/staking/source/local"
	remotestakingsource "github.com/stalwart-algoritmiclab/callisto/modules/staking/source/remote"
	coresource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/core"
	remotecoreSource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/core/source/remote"
	exchangerSource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source"
	remoteexchangerSource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source/remote"
	exchangersource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source"
	remoteexchangersource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source/remote"
	faucetsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/faucet/source"
	remotefaucetsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/faucet/source/remote"
	coretypes "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/core"
	securedsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/secured/source"
	remotesecuredsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/secured/source/remote"
	exchangertypes "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
	faucettypes "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
	securedtypes "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
)

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source

	ExchangerSource exchangersource.Source
	FaucetSource    faucetsource.Source
	SecuredSource   securedsource.Source
	CoreSource      coresource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	app := simapp.NewSimApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, nil, nil,
	)

	sources := &Sources{
		BankSource: localbanksource.NewSource(source, banktypes.QueryServer(app.BankKeeper)),
		// DistrSource:    localdistrsource.NewSource(source, distrtypes.QueryServer(app.DistrKeeper)),
		GovSource:      localgovsource.NewSource(source, govtypesv1.QueryServer(app.GovKeeper)),
		MintSource:     localmintsource.NewSource(source, minttypes.QueryServer(app.MintKeeper)),
		SlashingSource: localslashingsource.NewSource(source, slashingtypes.QueryServer(app.SlashingKeeper)),
		StakingSource:  localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: app.StakingKeeper}),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(app, "keys")
	if err != nil {
		return nil, err
	}

	err = source.MountTransientStores(app, "tkeys")
	if err != nil {
		return nil, err
	}

	err = source.MountMemoryStores(app, "memKeys")
	if err != nil {
		return nil, err
	}

	err = source.InitStores()
	if err != nil {
		return nil, err
	}

	return sources, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:     remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
		DistrSource:    remotedistrsource.NewSource(source, distrtypes.NewQueryClient(source.GrpcConn)),
		GovSource:      remotegovsource.NewSource(source, govtypesv1.NewQueryClient(source.GrpcConn)),
		MintSource:     remotemintsource.NewSource(source, minttypes.NewQueryClient(source.GrpcConn)),
		SlashingSource: remoteslashingsource.NewSource(source, slashingtypes.NewQueryClient(source.GrpcConn)),
		StakingSource:  remotestakingsource.NewSource(source, stakingtypes.NewQueryClient(source.GrpcConn)),

		// Custom stwart modules
		FaucetSource:    remotefaucetsource.NewSource(source, faucettypes.NewQueryClient(source.GrpcConn)),
		ExchangerSource: remoteexchangersource.NewSource(source, exchangertypes.NewQueryClient(source.GrpcConn)),
		SecuredSource:   remotesecuredsource.NewSource(source, securedtypes.NewQueryClient(source.GrpcConn)),
		CoreSource:      remotecoreSource.NewSource(source, coretypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
