/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

import (
	"fmt"

	"github.com/forbole/juno/v6/node/remote"

	"github.com/cosmos/cosmos-sdk/codec"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/forbole/juno/v6/node/local"

	nodeconfig "github.com/forbole/juno/v6/node/config"

	banksource "github.com/stalwart-algoritmiclab/callisto/modules/bank/source"
	localbanksource "github.com/stalwart-algoritmiclab/callisto/modules/bank/source/local"
	remotebanksource "github.com/stalwart-algoritmiclab/callisto/modules/bank/source/remote"
	distrsource "github.com/stalwart-algoritmiclab/callisto/modules/distribution/source"
	localdistrsource "github.com/stalwart-algoritmiclab/callisto/modules/distribution/source/local"
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
	"github.com/stalwart-algoritmiclab/callisto/utils/simapp"
	exchangerSource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source"
	remoteexchangerSource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/exchanger/source/remote"
	faucetsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/faucet/source"
	remotefaucetsource "github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/faucet/source/remote"
	exchangertypes "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
	faucettypes "github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
)

type Sources struct {
	BankSource     banksource.Source
	DistrSource    distrsource.Source
	GovSource      govsource.Source
	MintSource     mintsource.Source
	SlashingSource slashingsource.Source
	StakingSource  stakingsource.Source

	ExchangerSource exchangerSource.Source
	FaucetSource    faucetsource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, cdc codec.Codec) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, cdc)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, cdc codec.Codec) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, cdc)
	if err != nil {
		return nil, err
	}

	app := simapp.NewSimApp(cdc)

	sources := &Sources{
		BankSource:     localbanksource.NewSource(source, banktypes.QueryServer(app.BankKeeper)),
		DistrSource:    localdistrsource.NewSource(source, distrkeeper.NewQuerier(app.DistrKeeper)),
		GovSource:      localgovsource.NewSource(source, govkeeper.NewQueryServer(&app.GovKeeper)),
		MintSource:     localmintsource.NewSource(source, mintkeeper.NewQueryServerImpl(app.MintKeeper)),
		SlashingSource: localslashingsource.NewSource(source, slashingtypes.QueryServer(app.SlashingKeeper)),
		StakingSource:  localstakingsource.NewSource(source, stakingkeeper.Querier{Keeper: app.StakingKeeper}),
	}

	// Mount and initialize the stores
	err = source.MountKVStores(app, "keys")
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
		ExchangerSource: remoteexchangerSource.NewSource(source, exchangertypes.NewQueryClient(source.GrpcConn)),
	}, nil
}
