/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package actions

import (
	"github.com/forbole/juno/v5/node/remote"
	"gopkg.in/yaml.v3"
)

// Config contains the configuration about the actions module
type Config struct {
	Host string          `yaml:"host"`
	Port uint            `yaml:"port"`
	Node *remote.Details `yaml:"node,omitempty"`
}

// NewConfig returns a new Config instance
func NewConfig(host string, port uint, remoteDetails *remote.Details) *Config {
	return &Config{
		Host: host,
		Port: port,
		Node: remoteDetails,
	}
}

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Host: "127.0.0.1",
		Port: 3000,
		Node: nil,
	}
}

func ParseConfig(bz []byte) (*Config, error) {
	type T struct {
		Config *Config `yaml:"actions"`
	}
	var cfg T
	err := yaml.Unmarshal(bz, &cfg)

	if cfg.Config == nil {
		return DefaultConfig(), nil
	}

	return cfg.Config, err
}
