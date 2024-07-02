/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package staking

import (
	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewStakingCmd returns the Cobra command that allows to fix all the things related to the x/staking module
func NewStakingCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking",
		Short: "Fix things related to the x/staking module",
	}

	cmd.AddCommand(
		poolCmd(parseConfig),
		validatorsCmd(parseConfig),
	)

	return cmd
}
