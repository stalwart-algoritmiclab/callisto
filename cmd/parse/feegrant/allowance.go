/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feegrant

import (
	"encoding/hex"
	"fmt"
	"sort"

	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/forbole/juno/v6/types/config"

	"github.com/stalwart-algoritmiclab/callisto/modules/feegrant"
	"github.com/stalwart-algoritmiclab/callisto/utils"

	"github.com/spf13/cobra"

	"github.com/stalwart-algoritmiclab/callisto/database"

	tmctypes "github.com/cometbft/cometbft/rpc/core/types"

	feegranttypes "cosmossdk.io/x/feegrant"
	"github.com/rs/zerolog/log"
)

// allowanceCmd returns the Cobra command allowing to fix all things related to fee grant allowance
func allowanceCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "allowance",
		Short: "Fix granted and revoked allowances to the latest height",
		RunE: func(cmd *cobra.Command, args []string) error {
			parseCtx, err := parsecmdtypes.GetParserContext(config.Cfg, parseConfig)
			if err != nil {
				return err
			}

			// Get the database
			db := database.Cast(parseCtx.Database)

			// Build feegrant module
			feegrantModule := feegrant.NewModule(utils.GetCodec(), db)

			// Get the accounts
			// Collect all the transactions
			var txs []*tmctypes.ResultTx

			// Get all the MsgGrantAllowance txs
			query := fmt.Sprintf("message.action='%s'", feegranttypes.EventTypeSetFeeGrant)
			grantAllowanceTxs, err := utils.QueryTxs(parseCtx.Node, query)
			if err != nil {
				return err
			}
			txs = append(txs, grantAllowanceTxs...)

			// Get all the MsgRevokeAllowance txs
			query = fmt.Sprintf("message.action='%s'", feegranttypes.EventTypeRevokeFeeGrant)
			revokeAllowanceTxs, err := utils.QueryTxs(parseCtx.Node, query)
			if err != nil {
				return err
			}
			txs = append(txs, revokeAllowanceTxs...)

			// Sort the txs based on their ascending height
			sort.Slice(txs, func(i, j int) bool {
				return txs[i].Height < txs[j].Height
			})

			for _, tx := range txs {
				log.Debug().Int64("height", tx.Height).Msg("parsing transaction")
				transaction, err := parseCtx.Node.Tx(hex.EncodeToString(tx.Tx.Hash()))
				if err != nil {
					return err
				}

				// Handle only the MsgGrantAllowance and MsgRevokeAllowance instances
				for index, sdkMsg := range transaction.GetMsgs() {
					_, isMsgGrantAllowance := sdkMsg.(*feegranttypes.MsgGrantAllowance)
					_, isMsgRevokeAllowance := sdkMsg.(*feegranttypes.MsgRevokeAllowance)

					if !isMsgGrantAllowance && !isMsgRevokeAllowance {
						continue
					}

					err = feegrantModule.HandleMsg(index, transaction.Body.Messages[index], transaction)
					if err != nil {
						return fmt.Errorf("error while handling feegrant module message: %s", err)
					}
				}
			}

			return nil
		},
	}
}
