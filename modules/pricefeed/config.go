/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package pricefeed

import (
	"gopkg.in/yaml.v3"

	"github.com/stalwart-algoritmiclab/callisto/types"
)

// Config contains the configuration about the pricefeed module
type Config struct {
	Tokens []types.Token `yaml:"tokens"`
}

var PricefeedCfg *Config

// NewConfig returns a new Config instance
func NewConfig(tokens []types.Token) *Config {
	return &Config{
		Tokens: tokens,
	}
}

func ParseConfig(bz []byte) (*Config, error) {
	type T struct {
		Config *Config `yaml:"pricefeed"`
	}
	var cfg T
	err := yaml.Unmarshal(bz, &cfg)
	PricefeedCfg = cfg.Config

	return cfg.Config, err
}
