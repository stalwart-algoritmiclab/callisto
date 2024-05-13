/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package distribution

import (
	"github.com/cosmos/cosmos-sdk/codec"

	distrsource "github.com/forbole/callisto/modules/distribution/source"

	"github.com/forbole/juno/v6/modules"

	"github.com/forbole/callisto/database"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.GenesisModule            = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
	_ modules.MessageModule            = &Module{}
	_ modules.AuthzMessageModule       = &Module{}
)

// Module represents the x/distr module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	source distrsource.Source
}

// NewModule returns a new Module instance
func NewModule(source distrsource.Source, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		source: source,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "distribution"
}
