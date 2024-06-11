/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package testutils

import (
	"time"
)

func NewDurationPointer(duration time.Duration) *time.Duration {
	return &duration
}

func NewTimePointer(time time.Time) *time.Time {
	return &time
}
