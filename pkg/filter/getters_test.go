/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package filter

import (
	"reflect"
	"testing"
)

func TestFilter_Getters(t *testing.T) {
	tests := []struct {
		name    string
		value   any
		want    any
		getFunc func(f Filter, field string) any
	}{
		{
			name:  "get int from string",
			value: "10",
			want:  10,
			getFunc: func(f Filter, field string) any {
				return f.GetInt(field)
			},
		},
		{
			name:  "get int from int",
			value: 10,
			want:  10,
			getFunc: func(f Filter, field string) any {
				return f.GetInt(field)
			},
		},
		{
			name:  "get int from invalid string",
			value: "invalid",
			want:  0,
			getFunc: func(f Filter, field string) any {
				return f.GetInt(field)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				testField = "test"
				filter    = NewFilter().SetArgument(testField, tt.value)
				got       = tt.getFunc(filter, testField)
			)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArgument() = %v, want %v", got, tt.want)
			}
		})
	}
}
