/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package stwart

// Errors messages
const (
	errorMsgUpdateClient = "error while unpacking message: no concrete type registered for type URL /ibc.core.client.v1.MsgUpdateClient against interface *types.Msg"
)

// handleErrors - handle errors
func (m *Module) handleErrors(err error) error {
	switch err.Error() {
	// TODO: check update bd-juno to cosmos-sdk v0.50+
	// https://github.com/forbole/callisto/issues/701
	case errorMsgUpdateClient:
		m.logger.Info("Skip", "error", err)
		return nil
	default:
		return err
	}
}
