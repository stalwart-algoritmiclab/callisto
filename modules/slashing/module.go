/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package slashing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v5/modules"

	"github.com/stalwart-algoritmiclab/callisto/database"
	slashingsource "github.com/stalwart-algoritmiclab/callisto/modules/slashing/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.BlockModule   = &Module{}
)

// Module represent x/slashing module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source slashingsource.Source
}

// NewModule returns a new Module instance
func NewModule(source slashingsource.Source, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "slashing"
}
