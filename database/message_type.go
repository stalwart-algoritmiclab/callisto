/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database

import (
	types "github.com/forbole/callisto/types"
)

// SaveMessageType stores the given message type inside the database
func (db *Db) SaveMessageType(msg *types.MessageType) error {
	stmt := `
INSERT INTO message_type(type, module, label, height) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT (type) DO NOTHING`

	_, err := db.SQL.Exec(stmt, msg.Type, msg.Module, msg.Label, msg.Height)
	return err
}
