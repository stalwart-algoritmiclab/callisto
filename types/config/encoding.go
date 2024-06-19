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

	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/core"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/exchanger"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/faucet"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/feepolicy"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/rates"
	"github.com/stalwart-algoritmiclab/callisto/proto/stwartchain/secured"
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
		core.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		exchanger.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		faucet.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		rates.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		secured.RegisterInterfaces(encodingConfig.InterfaceRegistry)
		feepolicy.RegisterInterfaces(encodingConfig.InterfaceRegistry)
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
