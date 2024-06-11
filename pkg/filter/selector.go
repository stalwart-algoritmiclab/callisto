/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package filter

import (
	"fmt"
	"strconv"
	"strings"
)

// Clause enum.
const (
	clauseWhere = " WHERE "
	clauseAND   = " AND "
	noClause    = ""
)

type (
	// selector custom sql-builder for select queries.
	selector struct {
		builder *strings.Builder // builder for query.
		clause  string           // WHERE or AND clause.
		counter int              // arguments counter.
		args    []any            // all arguments used in the query.
	}
)

// newSelector constructor.
func newSelector() selector {
	return selector{
		builder: &strings.Builder{},
		counter: 1,
		clause:  clauseWhere,
	}
}

// GetInStatement helper function to get IN statement with specified field and values amount.
func (s *selector) GetInStatement(field string, amount int) string {
	placeHolders := make([]string, 0, amount)

	for i := 0; i < amount; i++ {
		placeHolders = append(placeHolders, fmt.Sprintf("$%d", s.counter))
		s.counter++
	}

	return fmt.Sprintf(field+" IN (%s)", strings.Join(placeHolders, ","))
}

// SetArguments setup where clause with field and their values.
//
//	condition defines which operator is used (WHERE or AND).
func (s *selector) SetArguments(arguments map[string][]interface{}, condition string) {
	s.setArguments(arguments, condition)
	s.clause = clauseAND
}

// setArguments setup field and their values.
//
//	condition defines which operator is used (WHERE or AND).
func (s *selector) setArguments(arguments map[string][]interface{}, condition string) {
	subQuery := make([]string, 0, len(arguments))

	for field, values := range arguments {
		subQuery = append(subQuery, s.GetInStatement(field, len(values)))
		s.args = append(s.args, values...)
	}

	_, _ = fmt.Fprintf(s.builder, "%s(%s)", s.clause, strings.Join(subQuery, " "+condition+" "))
}

// SetArgumentsGroups setups where clause with multiple groups of arguments.
func (s *selector) SetArgumentsGroups(groups []argumentsGroup) {
	_, _ = fmt.Fprintf(s.builder, "%s", s.clause)
	s.clause = noClause // reset clause.

	for _, group := range groups {
		if len(group.leftCondition) > 0 {
			_, _ = fmt.Fprintf(s.builder, " %s ", group.leftCondition)
		}

		s.setArguments(group.argumentsList, group.innerCondition.String())
	}

	s.clause = clauseAND
}

// SetClause setups where clause .
func (s *selector) SetClause(clause string) {
	s.clause = clause
}

// SetPredicates setup where clause with predicates.
// Condition defines which operator will be used (WHERE or AND).
func (s *selector) SetPredicates(predicates []predicator, condition string) {
	var (
		subQuery string
		args     []any
		query    = make([]string, 0, len(predicates))
	)

	for _, pr := range predicates {
		subQuery, args, s.counter = pr.build(condition, s.counter)
		query = append(query, subQuery)
		s.args = append(s.args, args...)
	}

	s.builder.WriteString(s.clause + strings.Join(query, " "+condition+" "))
	s.clause = clauseAND
}

// SetSearch setups LIKE operator with multiple fields in the query.
//
//	Example: name LIKE Art% OR surname LIKE %Art
func (s *selector) SetSearch(searcher searcher) {
	subQuery := make([]string, 0, len(searcher.fields))

	for _, field := range searcher.fields {
		subQuery = append(subQuery, fmt.Sprintf("LOWER(%s::text) LIKE LOWER($%d)", field, s.counter))
	}

	_, _ = fmt.Fprintf(s.builder, "%s(%s)", s.clause, strings.Join(subQuery, " "+ConditionOR.String()+" "))

	s.args = append(s.args, searcher.pattern)
	s.counter++
	s.clause = clauseAND
}

// PrepareSelector prepares default select from query.
func (s selector) PrepareSelector(table string, distinction bool, fields ...string) {
	if len(fields) == 0 {
		fields = append(fields, "*")
	}

	var args string
	if distinction {
		args = "DISTINCT "
	}

	_, _ = fmt.Fprintf(s.builder, "SELECT %s FROM %s", args+strings.Join(fields, ","), table)
}

// SetLimit setups limit in the query.
//
//	Example: LIMIT 10
func (s selector) SetLimit(limit uint64) {
	s.builder.WriteString(" LIMIT " + strconv.FormatUint(limit, 10))
}

// SetOffset setups offset in the query.
//
//	Example: OFFSET 10
func (s selector) SetOffset(offset uint64) {
	s.builder.WriteString(" OFFSET " + strconv.FormatUint(offset, 10))
}

// SetGroupBy setups group by statement in the query.
//
//	Example: GROUP BY status
func (s selector) SetGroupBy(fields ...string) {
	s.builder.WriteString(" GROUP BY " + strings.Join(fields, ", "))
}

// SetSorting setups order by statement in the query.
//
//	Example: ORDER BY id ASC, status DESC
func (s selector) SetSorting(sortArguments map[string]SortDirection) {
	subQuery := make([]string, 0, len(sortArguments))

	for field, direction := range sortArguments {
		subQuery = append(subQuery, field+" "+string(direction))
	}

	s.builder.WriteString(" ORDER BY " + strings.Join(subQuery, ", "))
}

// SetSuffix writes any text after FROM operator.
// Can be used to make JOIN statements.
func (s selector) SetSuffix(suffix string) {
	s.builder.WriteString(suffix)
}

// Build returns query and arguments from query-builder
func (s selector) Build() (query string, args []any) {
	return s.builder.String(), s.args
}
