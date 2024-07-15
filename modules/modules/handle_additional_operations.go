/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package modules

// RunAdditionalOperations implements AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	return m.db.InsertEnableModules(m.cfg.Modules)
}
