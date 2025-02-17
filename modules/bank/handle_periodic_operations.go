/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package bank

import (
	"fmt"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/modules/utils"
)

// RegisterPeriodicOperations implements modules.Module
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "bank").Msg("setting up periodic tasks")

	if _, err := scheduler.Every(10).Minutes().Do(func() {
		utils.WatchMethod(m.UpdateSupply)
	}); err != nil {
		return fmt.Errorf("error while setting up bank periodic operation: %s", err)
	}

	return nil
}

// UpdateSupply updates the supply of all the tokens
func (m *Module) UpdateSupply() error {
	log.Trace().Str("module", "bank").Str("operation", "total supply").
		Msg("updating total supply")

	block, err := m.db.GetLastBlockHeightAndTimestamp()
	if err != nil {
		return fmt.Errorf("error while getting latest block height: %s", err)
	}

	supply, err := m.keeper.GetSupply(block.Height)
	if err != nil {
		return err
	}

	return m.db.SaveSupply(supply, block.Height)
}
