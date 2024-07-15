/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	utils "github.com/forbole/juno/v6/node/local"
	txtypes "github.com/forbole/juno/v6/types"

	"github.com/stalwart-algoritmiclab/callisto/database/types"
	"github.com/stalwart-algoritmiclab/callisto/pkg/errs"
	"github.com/stalwart-algoritmiclab/callisto/pkg/filter"
)

// GetTransaction - get transaction from database
func (db *Db) GetTransaction(filter filter.Filter) (*txtypes.Transaction, error) {
	query, args := filter.SetLimit(1).Build("transaction")
	var result types.TransactionRow
	if err := db.Sqlx.Get(&result, query, args...); err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return &txtypes.Transaction{}, errs.Internal{Cause: err.Error()}
		}

		return &txtypes.Transaction{}, errs.NotFound{What: "transaction"}
	}

	return db.toTxTypesTx(result)
}

// GetTransactions - get transactions from database
func (db *Db) GetTransactions(filter filter.Filter) ([]*txtypes.Transaction, error) {
	query, args := filter.Build("transaction")

	var result []types.TransactionRow
	if err := db.Sqlx.Select(&result, query, args...); err != nil {
		return []*txtypes.Transaction{}, errs.Internal{Cause: err.Error()}
	}

	if len(result) == 0 {
		return []*txtypes.Transaction{}, errs.NotFound{What: "transaction"}
	}

	transactions := make([]*txtypes.Transaction, 0, len(result))
	for _, tx := range result {
		transaction, err := db.toTxTypesTx(tx)
		if err != nil {
			return []*txtypes.Transaction{}, errs.Internal{Cause: err.Error()}
		}

		transactions = append(transactions, transaction)
	}

	for _, tx := range transactions {

		_ = test(*tx)
	}

	return transactions, nil
}

func test(transaction txtypes.Transaction) error {

	return nil
}

// toTxTypesTx - convert database row to Tx
func (db *Db) toTxTypesTx(tx types.TransactionRow) (*txtypes.Transaction, error) {
	var err error
	result, err := txtypes.NewTransaction(
		utils.NewTxResponseFromSdkTxResponse(&sdk.TxResponse{}, &txtypes.Tx{}),
		newTxFromSdkTx(&sdktx.Tx{}, &sdktx.TxBody{}, &txtypes.AuthInfo{}),
	)
	if err != nil {
		return nil, err
	}

	if !tx.Success {
		result.TxResponse.Code = 1
	}
	var anyRaw []json.RawMessage
	if err = json.Unmarshal(tx.Messages, &anyRaw); err != nil {
		return nil, err
	}

	for _, raw := range anyRaw {
		msg := txtypes.StandardMessage{}
		if err := json.Unmarshal(raw, &msg); err != nil {
			return nil, fmt.Errorf("failed to unmarshal StandardMessage: %v", err)
		}

		result.Body.Messages = append(result.Body.Messages, &msg)
	}

	result.Signatures = make([][]byte, len(tx.Signatures))
	for index, sig := range tx.Signatures {
		if result.Signatures[index], err = base64.StdEncoding.DecodeString(sig); err != nil {
			return nil, err
		}
	}

	var sigInfoRaw []json.RawMessage
	if err = json.Unmarshal(tx.SignerInfos, &sigInfoRaw); err != nil {
		return nil, err
	}

	for _, sig := range sigInfoRaw {
		sigInfo := txtypes.SignerInfo{}
		if err = json.Unmarshal(sig, &sigInfo); err != nil {
			return nil, err
		}
		result.AuthInfo.SignerInfos = append(result.AuthInfo.SignerInfos, &sigInfo)
	}

	result.AuthInfo.Fee = &txtypes.Fee{}
	if err = json.Unmarshal(tx.Fee, result.AuthInfo.Fee); err != nil {
		return nil, err
	}

	result.Logs = sdk.ABCIMessageLogs{}
	if err = json.Unmarshal(tx.Logs, &result.Logs); err != nil {
		return nil, err
	}

	block, err := db.GetBlock(filter.NewFilter().SetArgument(types.FieldHeight, tx.Height))
	if err != nil {
		return nil, err
	}

	result.TxHash = tx.Hash
	result.Height = uint64(tx.Height)
	result.Body.Memo = tx.Memo
	result.GasWanted = uint64(tx.GasWanted)
	result.GasUsed = uint64(tx.GasUsed)
	result.RawLog = tx.RawLog
	result.Timestamp = block.Timestamp.Format(time.RFC3339)

	return result, nil
}

// newTxFromSdkTx allows to build a new Tx instance from the given tx.Tx
func newTxFromSdkTx(tx *sdktx.Tx, body *sdktx.TxBody, authInfo *txtypes.AuthInfo) *txtypes.Tx {
	return &txtypes.Tx{
		Tx:       tx,
		Body:     newTxBodFromSdkTxBody(body),
		AuthInfo: authInfo,
	}
}

// newTxBodFromSdkTxBody allows to build a new TxBody instance from the given tx.TxBody
func newTxBodFromSdkTxBody(body *sdktx.TxBody) *txtypes.TxBody {
	return &txtypes.TxBody{
		TxBody:        body,
		TimeoutHeight: body.TimeoutHeight,
		Messages:      []txtypes.Message{},
	}
}
