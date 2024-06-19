/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package modules

// RunAdditionalOperations implements AdditionalOperationsModule
func (m *Module) RunAdditionalOperations() error {
	err := m.db.InsertEnableModuleTypes()
	if err != nil {
		return err
	}

	return m.db.InsertEnableModules(m.cfg.Modules)
}
