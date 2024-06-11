/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

type SoftwareUpgradePlanRow struct {
	ProposalID    uint64 `db:"proposal_id"`
	PlanName      string `db:"plan_name"`
	UpgradeHeight int64  `db:"upgrade_height"`
	Info          string `db:"info"`
	Height        int64  `db:"height"`
}

func NewSoftwareUpgradePlanRow(
	proposalID uint64, planName string, upgradeHeight int64, info string, height int64,
) SoftwareUpgradePlanRow {
	return SoftwareUpgradePlanRow{
		ProposalID:    proposalID,
		PlanName:      planName,
		UpgradeHeight: upgradeHeight,
		Info:          info,
		Height:        height,
	}
}
