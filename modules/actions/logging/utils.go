/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package logging

import (
	"fmt"
	"net/http"
	"time"
)

func SuccessCounter(path string) {
	ActionCounter.WithLabelValues(path, fmt.Sprintf("%d", http.StatusOK)).Inc()
}

func ErrorCounter(path string) {
	ActionErrorCounter.WithLabelValues(path, fmt.Sprintf("%d", http.StatusInternalServerError)).Inc()
}

func ReponseTimeBuckets(path string, start time.Time) {
	ActionResponseTime.WithLabelValues(path).
		Observe(time.Since(start).Seconds())
}
