/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package exchanger

const (
	// ModuleName defines the module name
	ModuleName = "exchanger"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_exchanger"
)

var (
	ParamsKey = []byte("p_exchanger")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
