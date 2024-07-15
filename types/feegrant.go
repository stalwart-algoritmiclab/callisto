/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

import feegranttypes "cosmossdk.io/x/feegrant"

// FeeGrant represents the x/feegrant module
type FeeGrant struct {
	feegranttypes.Grant
	Height int64
}

// NewFeeGrant allows to build a new FeeGrant instance
func NewFeeGrant(feegrant feegranttypes.Grant, height int64) FeeGrant {
	return FeeGrant{
		feegrant,
		height,
	}
}

type GrantRemoval struct {
	Grantee string
	Granter string
	Height  int64
}

// NewGrantRemoval allows to build a new GrantRemoval instance
func NewGrantRemoval(grantee string, granter string, height int64) GrantRemoval {
	return GrantRemoval{
		grantee,
		granter,
		height,
	}
}
