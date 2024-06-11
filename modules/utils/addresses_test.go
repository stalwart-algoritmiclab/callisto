/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package utils_test

import (
	"testing"

	"github.com/stalwart-algoritmiclab/callisto/modules/utils"

	"github.com/stretchr/testify/require"
)

func TestFilterNonAccountAddresses(t *testing.T) {
	addresses := []string{
		"cosmos1hafptm4zxy5nw8rd2pxyg83c5ls2v62tstzuv2",
		"cosmosvaloper1hafptm4zxy5nw8rd2pxyg83c5ls2v62t4lkfqe",
	}

	filtered := utils.FilterNonAccountAddresses(addresses)
	require.Equal(t, []string{
		"cosmos1hafptm4zxy5nw8rd2pxyg83c5ls2v62tstzuv2",
	}, filtered)
}
