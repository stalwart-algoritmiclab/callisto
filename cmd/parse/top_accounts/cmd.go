/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package top_accounts

import (
	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewTopAccountsCmd returns the command that allows to parse the top accounts
func NewTopAccountsCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "top-accounts",
	}

	cmd.AddCommand(
		allCmd(parseConfig),
	)

	return cmd
}
