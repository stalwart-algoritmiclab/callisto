/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package filter

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

// GetArgument returns argument's first value.
func (f Filter) GetArgument(key string) any {
	if len(f.arguments[key]) > 0 {
		return f.arguments[key][0]
	}

	return nil
}

// GetArguments returns argument's values.
func (f Filter) GetArguments(key string) []any {
	return f.arguments[key]
}

// GetAllArguments returns all arguments.
func (f Filter) GetAllArguments() map[string][]any {
	return f.arguments
}

// GetString returns argument's first value as string.
func (f Filter) GetString(key string) string {
	return cast.ToString(f.GetArgument(key))
}

// GetBool returns argument's first value as bool.
func (f Filter) GetBool(key string) bool {
	return cast.ToBool(f.GetArgument(key))
}

// GetInt returns argument's first value as int.
func (f Filter) GetInt(key string) int {
	return cast.ToInt(f.GetArgument(key))
}

// GetInt8 returns argument's first value as int8.
func (f Filter) GetInt8(key string) int8 {
	return cast.ToInt8(f.GetArgument(key))
}

// GetInt16 returns argument's first value as int16.
func (f Filter) GetInt16(key string) int16 {
	return cast.ToInt16(f.GetArgument(key))
}

// GetInt32 returns argument's first value as int32.
func (f Filter) GetInt32(key string) int32 {
	return cast.ToInt32(f.GetArgument(key))
}

// GetInt64 returns argument's first value as int64.
func (f Filter) GetInt64(key string) int64 {
	return cast.ToInt64(f.GetArgument(key))
}

// GetUint returns argument's first value as uint.
func (f Filter) GetUint(key string) uint {
	return cast.ToUint(f.GetArgument(key))
}

// GetUint8 returns argument's first value as uint8.
func (f Filter) GetUint8(key string) uint8 {
	return cast.ToUint8(f.GetArgument(key))
}

// GetUint16 returns argument's first value as uint16.
func (f Filter) GetUint16(key string) uint16 {
	return cast.ToUint16(f.GetArgument(key))
}

// GetUint32 returns argument's first value as uint32.
func (f Filter) GetUint32(key string) uint32 {
	return cast.ToUint32(f.GetArgument(key))
}

// GetUint64 returns argument's first value as uint64.
func (f Filter) GetUint64(key string) uint64 {
	return cast.ToUint64(f.GetArgument(key))
}

// GetFloat32 returns argument's first value as float32.
func (f Filter) GetFloat32(key string) float32 {
	return cast.ToFloat32(f.GetArgument(key))
}

// GetFloat64 returns argument's first value as float64.
func (f Filter) GetFloat64(key string) float64 {
	return cast.ToFloat64(f.GetArgument(key))
}

// GetTime returns argument's first value as time.Time.
func (f Filter) GetTime(key string) time.Time {
	return cast.ToTime(f.GetArgument(key))
}

// GetDecimal returns argument's first value as decimal.Decimal.
func (f Filter) GetDecimal(key string) decimal.Decimal {
	return decimal.NewFromFloat(cast.ToFloat64(f.GetArgument(key)))
}

// GetDuration returns argument's first value as time.Duration.
func (f Filter) GetDuration(key string) time.Duration {
	return cast.ToDuration(f.GetArgument(key))
}

// GetIntSlice returns argument's first value as []int.
func (f Filter) GetIntSlice(key string) []int {
	return cast.ToIntSlice(f.GetArguments(key))
}

// GetStringSlice returns argument's first value as []string.
func (f Filter) GetStringSlice(key string) []string {
	return cast.ToStringSlice(f.GetArguments(key))
}

// GetStringMap returns argument's first value as map[string]interface{}.
func (f Filter) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(f.GetArgument(key))
}

// GetStringMapString returns argument's first value as map[string]string.
func (f Filter) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(f.GetArgument(key))
}

// GetStringMapStringSlice returns argument's first value as map[string][]string.
func (f Filter) GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(f.GetArguments(key))
}
