/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package errs

import (
	"testing"
)

func TestInternalSrv_Error(t *testing.T) {
	type fields struct {
		Cause string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "correct error",
			fields: fields{Cause: "test"},
			want:   "internal_server_error" + divider + "test",
		},

		{
			name:   "empty error",
			fields: fields{Cause: ""},
			want:   "internal_server_error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Internal{
				Cause: tt.fields.Cause,
			}

			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldsValidation_Error(t *testing.T) {
	type fields struct {
		FieldErrors []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "[success] correct error",
			fields: fields{
				FieldErrors: []string{
					"cyrillic_min::min_length_is::70", "polish_min::min_length_is::111", "german_min::min_length_is::120",
				},
			},
			want: "validation_error" + divider + "cyrillic_min::min_length_is::70,polish_min::min_length_is::111,german_min::min_length_is::120",
		},

		{
			name:   "[success] empty error with initial slice",
			fields: fields{FieldErrors: []string{}},
			want:   "validation_error",
		},

		{
			name:   "[success] empty error with nil slice",
			fields: fields{FieldErrors: nil},
			want:   "validation_error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := FieldsValidation{
				Errors: tt.fields.FieldErrors,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("FieldsValidation.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
