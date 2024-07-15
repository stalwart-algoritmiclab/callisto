/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package auth

import (
	parsecmdtypes "github.com/forbole/juno/v6/cmd/parse/types"
	"github.com/spf13/cobra"
)

// NewAuthCmd returns the Cobra command that allows to fix all the things related to the x/auth module
func NewAuthCmd(parseCfg *parsecmdtypes.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Fix things related to the x/auth module",
	}

	cmd.AddCommand(
		vestingCmd(parseCfg),
	)

	return cmd
}
