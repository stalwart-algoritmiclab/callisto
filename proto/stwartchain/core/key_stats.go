/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package core

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StatsKeyPrefix is the prefix to retrieve all Stats
	StatsKeyPrefix = "Stats/value/"
)

// StatsKey returns the store key to retrieve a Stats from the index fields
func StatsKey(
	date string,
) []byte {
	var key []byte

	dateBytes := []byte(date)
	key = append(key, dateBytes...)
	key = append(key, []byte("/")...)

	return key
}
