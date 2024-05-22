package gov

import (
	"github.com/stalwart-algoritmiclab/callisto/types"
)

type DistrModule interface {
	UpdateParams(height int64) error
}

type MintModule interface {
	UpdateParams(height int64) error
	UpdateInflation() error
}

type SlashingModule interface {
	UpdateParams(height int64) error
}

type StakingModule interface {
	GetStakingPoolSnapshot(height int64) (*types.PoolSnapshot, error)
	UpdateParams(height int64) error
}
