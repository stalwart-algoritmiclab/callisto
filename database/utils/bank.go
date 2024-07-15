/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package utils

import "github.com/stalwart-algoritmiclab/callisto/types"

const (
	maxPostgreSQLParams = 65535
)

func SplitAccounts(accounts []types.Account, paramsNumber int) [][]types.Account {
	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	slices := make([][]types.Account, len(accounts)/maxBalancesPerSlice+1)

	sliceIndex := 0
	for index, account := range accounts {
		slices[sliceIndex] = append(slices[sliceIndex], account)

		if index > 0 && index%(maxBalancesPerSlice-1) == 0 {
			sliceIndex++
		}
	}

	return slices
}

// SplitTopAccounts splits the given top accounts into slices
func SplitTopAccounts(accounts []types.TopAccount, paramsNumber int) [][]types.TopAccount {
	if paramsNumber == 0 {
		return nil
	}

	maxBalancesPerSlice := maxPostgreSQLParams / paramsNumber
	if maxBalancesPerSlice == 0 {
		return nil
	}

	numSlices := (len(accounts) + maxBalancesPerSlice - 1) / maxBalancesPerSlice
	slices := make([][]types.TopAccount, numSlices)

	for i, account := range accounts {
		sliceIndex := i / maxBalancesPerSlice
		slices[sliceIndex] = append(slices[sliceIndex], account)
	}

	return slices
}
