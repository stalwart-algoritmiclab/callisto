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

type (
	// predicator defines contract for predicates.
	predicator interface {
		build(condition string, counter int) (query string, args []interface{}, currentCounter int)
		updateFields(updateFunc) predicator
	}

	// updateFunc updates field in user required way.
	updateFunc func(field string) string
)

// Operators for predicates enum.
const (
	operatorLt         = "<"
	operatorLtOrEq     = "<="
	operatorGt         = ">"
	operatorGtOrEq     = ">="
	operatorIsNull     = "IS NULL"
	operatorIsNotNull  = "IS NOT NULL"
	operatorBetween    = "BETWEEN"
	operatorNotBetween = "NOT BETWEEN"
)

type (
	// comparator stores key/value pairs for predicates.
	comparator map[string]interface{}

	// PredicateLt describes 'less then' modifier.
	PredicateLt comparator

	// PredicateGt describes 'greater then' modifier.
	PredicateGt comparator

	// PredicateLtOrEq describes 'less then or equal' modifier.
	PredicateLtOrEq comparator

	// PredicateGtOrEq describes 'greater than or equal' modifier.
	PredicateGtOrEq comparator
)

// build - returns sql-like string representation.
func (c comparator) build(operator string, condition string, counter int) (query string, args []interface{}, currentCounter int) {
	statements := make([]string, 0, len(c))

	for key, value := range c {
		statements = append(statements, fmt.Sprintf("%s %s $%d", key, operator, counter))
		args = append(args, value)
		counter++
	}

	return strings.Join(statements, " "+condition+" "), args, counter
}

// updateFields - updates fields with user specified function.
func (c comparator) updateFields(update updateFunc) comparator {
	args := make(comparator, len(c))

	for field, values := range c {
		args[update(field)] = values
	}

	return args
}

// build - returns sql-like string representation.
func (lt PredicateLt) build(condition string, counter int) (query string, args []interface{}, currentCounter int) {
	return comparator(lt).build(operatorLt, condition, counter)
}

// updateFields - updates fields with user specified function.
func (lt PredicateLt) updateFields(update updateFunc) predicator {
	return PredicateLt(comparator(lt).updateFields(update))
}

// build - returns sql-like string representation.
func (gt PredicateGt) build(condition string, counter int) (query string, args []interface{}, currentCounter int) {
	return comparator(gt).build(operatorGt, condition, counter)
}

// updateFields - updates fields with user specified function.
func (gt PredicateGt) updateFields(update updateFunc) predicator {
	return PredicateGt(comparator(gt).updateFields(update))
}

// build - returns sql-like string representation.
func (lt PredicateLtOrEq) build(condition string, counter int) (query string, args []interface{}, currentCounter int) {
	return comparator(lt).build(operatorLtOrEq, condition, counter)
}

// updateFields - updates fields with user specified function.
func (lt PredicateLtOrEq) updateFields(update updateFunc) predicator {
	return PredicateLtOrEq(comparator(lt).updateFields(update))
}

// build - returns sql-like string representation.
func (gt PredicateGtOrEq) build(condition string, counter int) (query string, args []interface{}, currentCounter int) {
	return comparator(gt).build(operatorGtOrEq, condition, counter)
}

// updateFields - updates fields with user specified function.
func (gt PredicateGtOrEq) updateFields(update updateFunc) predicator {
	return PredicateGtOrEq(comparator(gt).updateFields(update))
}

type (
	// predicateNullable stores fields for null check.
	predicateNullable struct {
		Fields   []string
		Operator string
	}
)

// PredicateIsNull describes 'is null' modifier.
func PredicateIsNull(fields ...string) predicator {
	return predicateNullable{
		Fields:   fields,
		Operator: operatorIsNull,
	}
}

// PredicateIsNotNull describes 'is not null' modifier.
func PredicateIsNotNull(fields ...string) predicator {
	return predicateNullable{
		Fields:   fields,
		Operator: operatorIsNotNull,
	}
}

// build - returns sql-like string representation.
func (p predicateNullable) build(condition string, counter int) (query string, args []interface{}, currentCounter int) {
	statements := make([]string, 0, len(p.Fields))

	for idx := range p.Fields {
		statements = append(statements, fmt.Sprintf("%s %s", p.Fields[idx], p.Operator))
	}

	return strings.Join(statements, " "+condition+" "), args, counter
}

// updateFields - updates fields with user specified function.
func (p predicateNullable) updateFields(update updateFunc) predicator {
	for idx := range p.Fields {
		p.Fields[idx] = update(p.Fields[idx])
	}

	return p
}

type (
	// predicateBetween describes 'between' modifier.
	predicateBetween struct {
		Field    string
		Operator string
		Min      interface{}
		Max      interface{}
	}
)

// PredicateBetween describes 'between' modifier.
func PredicateBetween(field string, min, max interface{}) predicator {
	return predicateBetween{
		Field:    field,
		Operator: operatorBetween,
		Min:      min,
		Max:      max,
	}
}

// PredicateNotBetween describes 'not between' modifier.
func PredicateNotBetween(field string, min, max interface{}) predicator {
	return predicateBetween{
		Field:    field,
		Operator: operatorNotBetween,
		Min:      min,
		Max:      max,
	}
}

// build - returns sql-like string representation.
func (p predicateBetween) build(_ string, counter int) (query string, args []interface{}, currentCounter int) {
	return fmt.Sprintf("%s %s $%d AND $%d", p.Field, p.Operator, counter, counter+1),
		[]interface{}{p.Min, p.Max},
		counter + 2
}

// updateFields - updates fields with user specified function.
func (p predicateBetween) updateFields(update updateFunc) predicator {
	p.Field = update(p.Field)
	return p
}

// predicateNotIn describes 'not in' modifier.
type predicateNotIn struct {
	Field  string
	Values []any
}

// PredicateNotIn describes 'not in' modifier.
func PredicateNotIn(field string, values ...any) predicator {
	return predicateNotIn{Field: field, Values: values}
}

// build - returns sql-like string representation.
func (p predicateNotIn) build(_ string, counter int) (query string, args []interface{}, currentCounter int) {
	statements := make([]string, 0, len(p.Values))

	for idx := range p.Values {
		statements = append(statements, fmt.Sprintf("$%d", counter))
		args = append(args, p.Values[idx])
		counter++
	}

	return fmt.Sprintf("%s NOT IN (%s)", p.Field, strings.Join(statements, ", ")), args, counter
}

// updateFields - updates fields with user specified function.
func (p predicateNotIn) updateFields(update updateFunc) predicator {
	p.Field = update(p.Field)
	return p
}
