package config

import (
	junoconf "github.com/forbole/juno/v5/types/config"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/forbole/callisto/v4/config/vault"
)

// CheckVaultConfig - checking the ability to load the configuration from vault
func CheckVaultConfig(serviceName string, cmdAndConf *cobra.Command) *cobra.Command {
	log.Info().Msg("using vault configuration")

	config, err := loadFromVault(serviceName)
	if err != nil {
		log.Err(err).Msg("failed to load config from vault")
		return cmdAndConf
	}

	cmdAndConf.PreRunE = func(_ *cobra.Command, _ []string) error {
		junoconf.Cfg = config
		return nil
	}

	return cmdAndConf
}

func loadFromVault(serviceName string) (junoconf.Config, error) {
	vaultClient, err := vault.NewClient(vault.NamespaceCubbyhole)
	if err != nil {
		return junoconf.Config{}, err
	}

	vaultData, err := vaultClient.Pull(serviceName)
	if err != nil {
		return junoconf.Config{}, err
	}

	return junoconf.DefaultConfigParser(vaultData)
}
