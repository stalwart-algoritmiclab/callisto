/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package feegrant

import (
	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewFeegrantCmd returns the Cobra command that allows to fix all the things related to the x/feegrant module
func NewFeegrantCmd(parseConfig *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "feegrant",
		Short: "Fix things related to the x/feegrant module",
	}

	cmd.AddCommand(
		allowanceCmd(parseConfig),
	)

	return cmd
}
