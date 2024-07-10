/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package top_accounts

import (
	"github.com/rs/zerolog/log"

	topaccounts "github.com/stalwart-algoritmiclab/callisto/modules/top_accounts"
)

// AddressQueue is a channel of addresses
type AddressQueue chan string

// NewQueue returns a new AddressQueue
func NewQueue(size int) AddressQueue {
	return make(chan string, size)
}

// Worker is the worker that will process the addresses
type Worker struct {
	queue             AddressQueue
	topaccountsModule *topaccounts.Module
}

// NewWorker returns a new Worker
func NewWorker(queue AddressQueue, topaccountsModule *topaccounts.Module) Worker {
	return Worker{
		queue:             queue,
		topaccountsModule: topaccountsModule,
	}
}

// start starts the worker
func (w Worker) start() {
	for address := range w.queue {
		err := w.topaccountsModule.RefreshAll(address)
		if err != nil {
			log.Error().Str("account", address).Err(err).Msg("re-enqueueing failed address")

			go func(address string) {
				w.queue <- address
			}(address)
		}

	}
}
