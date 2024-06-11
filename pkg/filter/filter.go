/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package filter

import (
	"reflect"
)

type (
	// argumentsList is a list of arguments for query.
	argumentsList map[string][]any

	// argumentsGroup is a group of arguments for query, with common innerCondition.
	argumentsGroup struct {
		argumentsList  argumentsList
		leftCondition  Condition
		innerCondition Condition
	}

	// Filter is a struct which contains user params to database.
	Filter struct {
		arguments       argumentsList
		argumentsGroups []argumentsGroup
		predicates      []predicator
		sorting         map[string]SortDirection
		search          searcher
		condition       Condition
		distinct        bool
		groupBy         []string
		limit           uint64
		offset          uint64
	}
)

// NewFilter constructor.
func NewFilter() Filter {
	return Filter{
		arguments:       make(map[string][]any),
		argumentsGroups: make([]argumentsGroup, 0),
		sorting:         make(map[string]SortDirection),
		condition:       ConditionAND, // default value
	}
}

// SetPredicate setups predicate statement in the filer.
func (f Filter) SetPredicate(predicate predicator) Filter {
	f.predicates = append(f.predicates, predicate)
	return f
}

// SetPredicates setups predicates statement in the filer.
func (f Filter) SetPredicates(predicates ...predicator) Filter {
	f.predicates = predicates
	return f
}

// SetArgumentsMap setups arguments map in the filter.
// Arguments map override current arguments.
//
//	Example: map[string][]any{
//					"name": {"Dima", "Michael"},
//					"surname": {"Surname"},
//			}
func (f Filter) SetArgumentsMap(arguments argumentsList) Filter {
	f.arguments = arguments
	return f
}

// SetArgument setups argument in the filter.
func (f Filter) SetArgument(field string, values ...any) Filter {
	f.arguments[field] = make([]any, 0, len(values))

	for _, value := range values {
		reflectValue := reflect.ValueOf(value)

		if isListType(value, reflectValue.Kind()) {
			for i := 0; i < reflectValue.Len(); i++ {
				f.arguments[field] = append(f.arguments[field], reflectValue.Index(i).Interface())
			}
		} else {
			f.arguments[field] = append(f.arguments[field], value)
		}
	}

	return f
}

// SetArgumentsGroup setups group of arguments with innerCondition in the filter.
func (f Filter) SetArgumentsGroup(condition Condition, argumentsList argumentsList, leftCondition ...Condition) Filter {
	leftCond := NoCondition
	if len(leftCondition) > 0 {
		leftCond = leftCondition[0]
	}

	f.argumentsGroups = append(f.argumentsGroups, argumentsGroup{
		leftCondition:  leftCond,
		argumentsList:  argumentsList,
		innerCondition: condition,
	})
	return f
}

// SetSort setups sorting by one field in the filter.
func (f Filter) SetSort(field string, direction SortDirection) Filter {
	f.sorting[field] = direction
	return f
}

// SetSortMap setups sorting by multiple fields in the filter.
// Sort map override current sort arguments.
//
//	Example: map[string]SortDirection {
//				"status": DirectionDescending,
//				"id": DirectionAscending,
//			}
func (f Filter) SetSortMap(sorting map[string]SortDirection) Filter {
	f.sorting = sorting
	return f
}

// SetDistinct setup distinct request in the filter.
// Distinct allows you to select only unique values.
func (f Filter) SetDistinct() Filter {
	f.distinct = true
	return f
}

// SetCondition setups linking condition for arguments in where clause.
//
//	condition can be AND or OR.
//	Example: status = new AND/OR id = 1.
func (f Filter) SetCondition(condition Condition) Filter {
	f.condition = condition
	return f
}

// SetLimit setups limit request in the filter.
// Limit defines how many entities to return in response.
func (f Filter) SetLimit(limit uint64) Filter {
	f.limit = limit
	return f
}

// SetOffset setups offset request in the filter.
// Offset defines how many entities to skip in the response.
func (f Filter) SetOffset(offset uint64) Filter {
	f.offset = offset
	return f
}

// SetGroupBy setups grouping by required field.
func (f Filter) SetGroupBy(fields ...string) Filter {
	f.groupBy = fields
	return f
}

// SetSearch setups search request in the filter.
func (f Filter) SetSearch(searchValue string, fields ...string) Filter {
	f.search.pattern = searchValue
	f.search.fields = fields
	return f
}

// Build returns prepared query and arguments for execution in database.
func (f Filter) Build(from string, fields ...string) (query string, args []any) {
	builder := newSelector()
	builder.PrepareSelector(from, f.distinct, fields...)

	return f.setupSelector(builder).Build()
}

// BuildCount returns prepared query and arguments for counting rows.
func (f Filter) BuildCount(from string) (query string, args []any) {
	return f.SetLimit(0).SetOffset(0).SetSortMap(nil).Build(from, queryCount)
}

// ExcludeArguments removes arguments from the filter.
func (f Filter) ExcludeArguments(keys ...string) Filter {
	filtered := make(map[string][]any)
	for k, v := range f.arguments {
		filtered[k] = v
	}

	for _, key := range keys {
		delete(filtered, key)
	}

	return f.SetArgumentsMap(filtered)
}

// ToJoiner converts filter to joiner model to work with JOIN queries.
func (f Filter) ToJoiner() Joiner {
	return NewJoiner(&f)
}

// setupSelector setups selector by parameters in the filter.
func (f Filter) setupSelector(builder selector) selector {
	if f.withArgumentsGroups() {
		builder.SetArgumentsGroups(f.argumentsGroups)
	}

	if f.withArguments() {
		builder.SetArguments(f.arguments, f.condition.String())
	}

	if f.withPredicates() {
		builder.SetPredicates(f.predicates, f.condition.String())
	}

	if f.withSearch() {
		builder.SetSearch(f.search)
	}

	if f.withGroupBy() {
		builder.SetGroupBy(f.groupBy...)
	}

	if f.withSorting() {
		builder.SetSorting(f.sorting)
	}

	if f.withLimit() {
		builder.SetLimit(f.limit)
	}

	if f.withOffset() {
		builder.SetOffset(f.offset)
	}

	return builder
}
