/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package rates

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v6/modules"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain/rates"
	"github.com/stalwart-algoritmiclab/callisto/modules/stwart/chain/rates/source"
)

var (
	_ modules.Module        = &Module{}
	_ modules.GenesisModule = &Module{}
	_ modules.MessageModule = &Module{}
)

// Module represents the x/rates module
type Module struct {
	cdc       codec.Codec
	db        *database.Db
	ratesRepo rates.Repository

	keeper source.Source
}

// NewModule returns a new Module instance
func NewModule(keeper source.Source, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		keeper:    keeper,
		cdc:       cdc,
		db:        db,
		ratesRepo: *rates.NewRepository(db.Sqlx, cdc),
	}
}

// Name implements modules.Module
func (m *Module) Name() string { return "stwart_rates" }
