/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package pricefeed

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v6/types/config"

	"github.com/stalwart-algoritmiclab/callisto/database"

	"github.com/forbole/juno/v6/modules"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
	_ modules.PeriodicOperationsModule   = &Module{}
)

// Module represents the module that allows to get the token prices
type Module struct {
	cfg *Config
	cdc codec.Codec
	db  *database.Db
}

// NewModule returns a new Module instance
func NewModule(cfg config.Config, cdc codec.Codec, db *database.Db) *Module {
	bz, err := cfg.GetBytes()
	if err != nil {
		panic(err)
	}

	pricefeedCfg, err := ParseConfig(bz)
	if err != nil {
		panic(err)
	}

	return &Module{
		cfg: pricefeedCfg,
		cdc: cdc,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pricefeed"
}
