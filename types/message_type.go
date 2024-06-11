/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package types

// MessageType represents the data of a single message type
type MessageType struct {
	Type   string
	Module string
	Label  string
	Height int64
}

// NewMessageType allows to build a new MessageType instance
func NewMessageType(msgType string, module string, label string, height int64) *MessageType {
	return &MessageType{
		Type:   msgType,
		Module: module,
		Label:  label,
		Height: height,
	}
}
