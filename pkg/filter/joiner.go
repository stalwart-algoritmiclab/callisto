/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package filter

import (
	"fmt"
	"strings"
)

// separator is a separator between table name and column name.
const separator = "."

type (
	// Joiner is a struct which make query with join to database.
	Joiner struct {
		selector   selector
		filter     *Filter
		tables     map[string][]string
		modifiers  map[string]string
		statements []string
	}
)

// NewJoiner constructor.
func NewJoiner(filter *Filter) Joiner {
	return Joiner{
		selector:  newSelector(),
		filter:    filter,
		tables:    make(map[string][]string),
		modifiers: make(map[string]string),
	}
}

// SetFieldModifier - prepare fields modifiers such function AS condition etc.
func (join Joiner) SetFieldModifier(field string, modifier string) Joiner {
	modifier = strings.Replace(modifier, field, "%s", 1)
	join.modifiers[field] = modifier
	return join
}

// PrepareTable adds a table with fields to joiner.
func (join Joiner) PrepareTable(tableName string, fields ...string) Joiner {
	join.tables[tableName] = fields
	return join
}

// PrepareJoinStatement prepares join query.
// 'joinStatement' specifies JOIN statement added to query.
// Table names must be identical to names provided in PrepareTable function.
// Multiple PrepareJoinStatement can be used.
//
//	Example: PrepareJoinStatement("INNER JOIN account on account.id = users.id")
func (join Joiner) PrepareJoinStatement(joinStatement string) Joiner {
	join.statements = append(join.statements, joinStatement)
	return join
}

// Build returns prepared query and arguments for execution in database.
// 'mainTable' specifies FROM which table to join.
//
//	Example: Build("users")
func (join Joiner) Build(mainTable string) (query string, args []interface{}) {
	var distinct string
	if join.filter.withDistinct() {
		distinct = "DISTINCT "
	}

	values := make([]string, 0, len(join.tables))

	for table, fields := range join.tables {
		for _, field := range fields {
			joinField := table + separator + field
			if mod, ok := join.modifiers[field]; ok {
				joinField = fmt.Sprintf(mod, joinField)
			}

			values = append(values, joinField)
		}
	}

	fmt.Fprintf(
		join.selector.builder,
		"SELECT %s FROM %s %s",
		distinct+strings.Join(values, ","),
		mainTable,
		strings.Join(join.statements, " "),
	)

	join.setupParameters()

	return join.filter.setupSelector(join.selector).Build()
}

// getTableByField returns table name by specified field.
func (join Joiner) getTableByField(field string) string {
	for table, fields := range join.tables {
		if isExist(fields, field) {
			return table
		}
	}

	return ""
}

// setupParameters adds table name to arguments in filter.
func (join Joiner) setupParameters() {
	if join.filter.withArguments() {
		args := make(map[string][]interface{}, len(join.filter.arguments))

		for field, values := range join.filter.arguments {
			args[join.prepareField(field)] = values
		}

		join.filter.arguments = args
	}

	if join.filter.withSearch() {
		for idx := range join.filter.search.fields {
			join.filter.search.fields[idx] = join.prepareField(join.filter.search.fields[idx])
		}
	}

	if join.filter.withSorting() {
		sort := make(map[string]SortDirection, len(join.filter.sorting))

		for field, value := range join.filter.sorting {
			sort[join.prepareField(field)] = value
		}

		join.filter.sorting = sort
	}

	if join.filter.withGroupBy() {
		for idx := range join.filter.groupBy {
			join.filter.groupBy[idx] = join.prepareField(join.filter.groupBy[idx])
		}
	}

	if join.filter.withPredicates() {
		for idx := range join.filter.predicates {
			join.filter.predicates[idx] = join.filter.predicates[idx].updateFields(join.prepareField)
		}
	}
}

// prepareField returns fields with table name if exists.
func (join Joiner) prepareField(field string) string {
	table := join.getTableByField(field)
	if table != "" {
		return table + separator + field
	} else {
		return field
	}
}
