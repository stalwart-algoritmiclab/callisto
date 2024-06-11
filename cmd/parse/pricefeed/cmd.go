/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package pricefeed

import (
	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewPricefeedCmd returns the Cobra command allowing to refresh pricefeed
func NewPricefeedCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pricefeed",
		Short: "Fix things related to the pricefeed module",
	}

	cmd.AddCommand(
		priceCmd(parseConfig),
		priceHistoryCmd(parseConfig),
	)

	return cmd
}
