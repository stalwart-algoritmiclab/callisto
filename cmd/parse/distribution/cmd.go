/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package distribution

import (
	parsecmdtypes "github.com/forbole/juno/v5/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewDistributionCmd returns the Cobra command allowing to fix various things related to the x/distribution module
func NewDistributionCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "distribution",
		Short: "Fix things related to the x/distribution module",
	}

	cmd.AddCommand(
		communityPoolCmd(parseConfig),
	)

	return cmd
}
