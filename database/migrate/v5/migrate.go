/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package v5

import (
	"fmt"

	"github.com/stalwart-algoritmiclab/callisto/modules/utils"
	"github.com/stalwart-algoritmiclab/callisto/types"
)

// Migrate implements database.Migrator
func (db *Migrator) Migrate() error {
	msgTypes, err := db.getMsgTypesFromMessageTable()
	if err != nil {
		return fmt.Errorf("error while getting message types rows: %s", err)
	}

	for _, msgType := range msgTypes {
		// migrate message types
		err = db.migrateMsgTypes(types.NewMessageType(
			msgType.Type,
			utils.GetModuleNameFromTypeURL(msgType.Type),
			utils.GetMsgFromTypeURL(msgType.Type),
			msgType.Height))

		if err != nil {
			return err
		}
	}
	return nil
}

// getMsgTypesFromMessageTable retrieves messages types stored in database inside message table
func (db *Migrator) getMsgTypesFromMessageTable() ([]MessageRow, error) {
	smt := "SELECT DISTINCT ON (type) type, transaction_hash, height FROM message ORDER BY type DESC"
	var rows []MessageRow
	err := db.SQL.Select(&rows, smt)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// migrateMsgTypes stores the given message type inside the database
func (db *Migrator) migrateMsgTypes(msg *types.MessageType) error {
	stmt := `
INSERT INTO message_type(type, module, label, height) 
VALUES ($1, $2, $3, $4) 
ON CONFLICT (type) DO NOTHING`

	_, err := db.SQL.Exec(stmt, msg.Type, msg.Module, msg.Label, msg.Height)
	return err
}
