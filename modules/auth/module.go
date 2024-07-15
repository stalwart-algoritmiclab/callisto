/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package auth

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v6/modules"
	"github.com/forbole/juno/v6/modules/messages"

	"github.com/stalwart-algoritmiclab/callisto/database"
	"github.com/stalwart-algoritmiclab/callisto/modules/auth/source"
)

var (
	_ modules.Module             = &Module{}
	_ modules.GenesisModule      = &Module{}
	_ modules.MessageModule      = &Module{}
	_ modules.AuthzMessageModule = &Module{}
)

// Module represents the x/auth module
type Module struct {
	cdc            codec.Codec
	db             *database.Db
	messagesParser messages.MessageAddressesParser
	source         source.Source
}

// NewModule builds a new Module instance
func NewModule(source source.Source, messagesParser messages.MessageAddressesParser, cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		messagesParser: messagesParser,
		source:         source,
		cdc:            cdc,
		db:             db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "auth"
}
