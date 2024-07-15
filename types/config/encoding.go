/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package config

import (
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/forbole/juno/v5/types/params"
	coretypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	exchangertypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
	faucettypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
	feepolicytypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"
	ratestypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"
	referralstypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"
	securedtypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"
)

// MakeEncodingConfig creates an EncodingConfig to properly handle all the messages
func MakeEncodingConfig(managers []module.BasicManager) func() params.EncodingConfig {
	return func() params.EncodingConfig {
		encodingConfig := params.MakeTestEncodingConfig()
		std.RegisterLegacyAminoCodec(encodingConfig.Amino)
		std.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		manager := mergeBasicManagers(managers)
		manager.RegisterLegacyAminoCodec(encodingConfig.Amino)
		manager.RegisterInterfaces(encodingConfig.InterfaceRegistry)

		// custom modules
		coretypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		exchangertypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		faucettypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		ratestypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		securedtypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		feepolicytypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		referralstypes.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		return encodingConfig
	}
}

// mergeBasicManagers merges the given managers into a single module.BasicManager
func mergeBasicManagers(managers []module.BasicManager) module.BasicManager {
	var union = module.BasicManager{}
	for _, manager := range managers {
		for k, v := range manager {
			union[k] = v
		}
	}
	return union
}
