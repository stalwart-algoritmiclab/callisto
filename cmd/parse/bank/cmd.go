/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package bank

import (
	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewBankCmd returns the Cobra command allowing to fix various things related to the x/bank module
func NewBankCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bank",
		Short: "Fix things related to the x/bank module",
	}

	cmd.AddCommand(
		supplyCmd(parseConfig),
	)

	return cmd
}
