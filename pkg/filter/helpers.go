/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package filter

import (
	"database/sql/driver"
	"reflect"
)

// withArguments checks if arguments field is specified.
func (f Filter) withArguments() bool { return len(f.arguments) != 0 }

// withArgumentsGroups checks if withArgumentsGroups field is not empty.
func (f Filter) withArgumentsGroups() bool { return len(f.argumentsGroups) != 0 }

// withSorting checks if sorting field is specified.
func (f Filter) withSorting() bool { return len(f.sorting) != 0 }

// withLimit checks if limit field is specified.
func (f Filter) withLimit() bool { return f.limit != 0 }

// withOffset checks if offset field is specified.
func (f Filter) withOffset() bool { return f.offset != 0 }

// withDistinct checks if distinct field is specified.
func (f Filter) withDistinct() bool { return f.distinct }

// withSearch checks if search field is specified.
func (f Filter) withSearch() bool { return f.search.pattern != "" }

// withPredicates checks if predicates field is specified.
func (f Filter) withPredicates() bool { return len(f.predicates) != 0 }

// withGroupBy checks if group by field is specified.
func (f Filter) withGroupBy() bool { return len(f.groupBy) != 0 }

// isListType checks if specified value is a list type.
func isListType(value interface{}, kind reflect.Kind) bool {
	if _, ok := value.(driver.Valuer); ok {
		return false
	}

	return kind == reflect.Array || kind == reflect.Slice
}

// isExist checks if specified field exists in array.
func isExist(array []string, value string) bool {
	for _, field := range array {
		if field == value {
			return true
		}
	}

	return false
}
