/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package daily_refetch

import (
	"github.com/forbole/juno/v6/node"

	callistodb "github.com/stalwart-algoritmiclab/callisto/database"

	"github.com/forbole/juno/v6/modules"
)

var (
	_ modules.Module                   = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

type Module struct {
	node     node.Node
	database *callistodb.Db
}

// NewModule builds a new Module instance
func NewModule(
	node node.Node,
	database *callistodb.Db,
) *Module {
	return &Module{
		node:     node,
		database: database,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "daily refetch"
}
