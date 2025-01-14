/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database

import (
	"fmt"
)

// InsertEnableModules allows to save enabled module into the database
func (db *Db) InsertEnableModules(modules []string) error {
	if len(modules) == 0 {
		return nil
	}

	// Remove existing modules
	stmt := "DELETE FROM modules WHERE TRUE"
	_, err := db.SQL.Exec(stmt)
	if err != nil {
		return fmt.Errorf("error while deleting modules: %s", err)
	}

	if len(modules) == 0 {
		return nil
	}

	var values []interface{}
	stmt = `INSERT INTO modules (module_name) VALUES`
	for key, value := range modules {
		stmt += fmt.Sprintf("($%d),", key+1)
		values = append(values, value)
	}
	stmt = stmt[:len(stmt)-1] // remove tailing ","
	stmt += " ON CONFLICT DO NOTHING"
	_, err = db.SQL.Exec(stmt, values...)
	if err != nil {
		return fmt.Errorf("error while storing modules: %s", err)
	}

	return nil
}
