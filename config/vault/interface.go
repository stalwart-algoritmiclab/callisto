/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package vault

// Vault interface.
type Vault interface {
	// Download data from vault via vault client to the file via path.
	Download(key, filepath string) error

	// Pull data from vault via vault client.
	Pull(key string) ([]byte, error)

	// PullRaw pull raw map data from vault via vault client.
	PullRaw(key string) (map[string]any, error)

	// Push data to the vault via vault client.
	Push(key string, value any) error

	// PushRaw data to the vault via vault client.
	PushRaw(key string, values map[string]any) error

	// Upload data from file to the vault.
	Upload(key, filepath string) error
}
