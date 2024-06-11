/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package v3

import (
	"fmt"
	"os"
	"path"

	"github.com/forbole/juno/v5/types/config"
	"gopkg.in/yaml.v3"
)

// GetConfig returns the configuration reading it from the config.yaml file present inside the home directory
func GetConfig() (Config, error) {
	file := path.Join(config.HomePath, "config.yaml")

	// Make sure the path exists
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("config file does not exist")
	}

	bz, err := os.ReadFile(file)
	if err != nil {
		return Config{}, fmt.Errorf("error while reading config files: %s", err)
	}

	var cfg Config
	err = yaml.Unmarshal(bz, &cfg)
	return cfg, err
}
