package filter

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestFilter_Build(t *testing.T) {
	uuidValue := uuid.New()

	type args struct {
		fields string
		table  string
	}
	tests := []struct {
		name      string
		filter    Filter
		args      args
		wantQuery string
		wantArgs  []any
	}{
		{
			name: "create query",
			filter: NewFilter().
				SetLimit(10).
				SetOffset(5).
				SetArgument("status", "active").
				SetSortMap(map[string]SortDirection{
					"user_id": DirectionAscending,
				}).
				SetSortMap(map[string]SortDirection{
					"id": DirectionDescending,
				}).
				SetGroupBy("status").
				SetSearch("Dima%", "name", "surname").
				SetDistinct(),
			args: args{
				fields: "id, user_id, name, surname, status, created_at, updated_at",
				table:  "users",
			},
			wantQuery: "SELECT DISTINCT id, user_id, name, surname, status, created_at, updated_at FROM users WHERE (status IN ($1)) AND (LOWER(name::text) LIKE LOWER($2) OR LOWER(surname::text) LIKE LOWER($2)) GROUP BY status ORDER BY id DESC LIMIT 10 OFFSET 5",
			wantArgs:  []any{"active", "Dima%"},
		},
		{
			name: "ArgumentsMap with Limit, Offset, SortMap, GroupBy, Search, Distinct create query",
			filter: NewFilter().
				SetLimit(10).
				SetOffset(5).
				SetArgumentsMap(map[string][]any{
					"name": {"Jon", "Michael"},
				}).
				SetSortMap(map[string]SortDirection{
					"user_id": DirectionAscending,
				}).
				SetGroupBy("status").
				SetSearch("Dima%", "name", "surname").
				SetDistinct(),
			args: args{
				fields: "id, user_id, name, surname, status, created_at, updated_at",
				table:  "users",
			},
			wantQuery: "SELECT DISTINCT id, user_id, name, surname, status, created_at, updated_at FROM users WHERE (name IN ($1,$2)) AND (LOWER(name::text) LIKE LOWER($3) OR LOWER(surname::text) LIKE LOWER($3)) GROUP BY status ORDER BY user_id ASC LIMIT 10 OFFSET 5",
			wantArgs:  []any{"Jon", "Michael", "Dima%"},
		},
		{
			name:   "select all",
			filter: NewFilter(),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users",
			wantArgs:  nil,
		},
		{
			name: "select all with Sort",
			filter: NewFilter().
				SetSort("user_id", DirectionAscending),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users ORDER BY user_id ASC",
			wantArgs:  nil,
		},
		{
			name: "select all with Argument and Condition",
			filter: NewFilter().
				SetArgument("status", "active").
				SetCondition(ConditionAND),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users WHERE (status IN ($1))",
			wantArgs:  []any{"active"},
		},
		{
			name: "select all with slice of arguments",
			filter: NewFilter().
				SetArgument("status", []string{"active", "blocked", "deleted"}).
				SetCondition(ConditionAND),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users WHERE (status IN ($1,$2,$3))",
			wantArgs:  []any{"active", "blocked", "deleted"},
		},
		{
			name: "select all with multiple arguments",
			filter: NewFilter().
				SetArgument("status", "active", "blocked", "deleted").
				SetCondition(ConditionAND),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users WHERE (status IN ($1,$2,$3))",
			wantArgs:  []any{"active", "blocked", "deleted"},
		},
		{
			name: "select all with uuid",
			filter: NewFilter().
				SetArgument("id", uuidValue).
				SetCondition(ConditionAND),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users WHERE (id IN ($1))",
			wantArgs:  []any{uuidValue},
		},
		{
			name: "select all with Argument Groups and Argument and Condition",
			filter: NewFilter().
				SetArgumentsGroup(ConditionOR, map[string][]any{
					"name": {"Alina"},
					"age":  {"18"},
				}).
				SetArgument("status", "active").
				SetCondition(ConditionAND),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users WHERE (name IN ($1) OR age IN ($2)) AND (status IN ($3))",
			wantArgs:  []any{"Alina", "18", "active"},
		},
		{
			name: "select all with 2 Argument Groups and Argument and Condition",
			filter: NewFilter().
				SetArgumentsGroup(ConditionOR, map[string][]any{
					"name": {"Alina"},
					"age":  {"18"},
				}).
				SetArgumentsGroup(ConditionOR, map[string][]any{
					"name": {"Cristina"},
					"age":  {"17"},
				},
					"LEFT_CONDITION", // left condition, used to join with previous argument group
				).
				SetArgument("status", "active").
				SetCondition(ConditionAND),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users WHERE (name IN ($1) OR age IN ($2)) LEFT_CONDITION (name IN ($3) OR age IN ($4)) AND (status IN ($5))",
			wantArgs:  []any{"Alina", "18", "Cristina", "17", "active"},
		},
		{
			name: "select all with 3 Argument Groups and Argument and Condition",
			filter: NewFilter().
				SetArgumentsGroup(ConditionOR, map[string][]any{
					"name": {"Alina"},
					"age":  {"18"},
				}).
				SetArgumentsGroup(ConditionOR, map[string][]any{
					"name": {"Cristina"},
					"age":  {"17"},
				},
					"LEFT_CONDITION_1", // left condition, used to join with previous argument group
				).
				SetArgumentsGroup(ConditionOR, map[string][]any{
					"name": {"Ira"},
					"age":  {"40"},
				},
					"LEFT_CONDITION_2", // left condition, used to join with previous argument group
				).
				SetArgument("status", "active").
				SetCondition(ConditionAND),
			args: args{
				table: "users",
			},
			wantQuery: "SELECT * FROM users WHERE (name IN ($1) OR age IN ($2)) LEFT_CONDITION_1 (name IN ($3) OR age IN ($4)) LEFT_CONDITION_2 (name IN ($5) OR age IN ($6)) AND (status IN ($7))",
			wantArgs:  []any{"Alina", "18", "Cristina", "17", "Ira", "40", "active"},
		},
		{
			name: "select all with Limit",
			filter: NewFilter().
				SetLimit(100),
			args: args{
				table: "requests",
			},
			wantQuery: `SELECT * FROM requests LIMIT 100`,
		},
		{
			name: "select all with Limit and Sort",
			filter: NewFilter().
				SetLimit(100).
				SetSort("id", DirectionDescending),
			args: args{
				table: "requests",
			},
			wantQuery: `SELECT * FROM requests ORDER BY id DESC LIMIT 100`,
		},
		{
			name: "select all with Limit, Sort, Argument",
			filter: NewFilter().
				SetLimit(100).
				SetSort("id", DirectionDescending).
				SetArgument("status", "active"),
			args: args{
				table: "requests",
			},
			wantQuery: `SELECT * FROM requests WHERE (status IN ($1)) ORDER BY id DESC LIMIT 100`,
			wantArgs:  []any{"active"},
		},
		{
			name: "select with Limit, Distinct",
			filter: NewFilter().
				SetLimit(100).
				SetDistinct(),
			args: args{
				fields: "id, user_id, first_name, last_name, address_main, address_secondary, city, zip_code, state, country, phone, email, comment, status, start_at, stop_at, created_at, updated_at, account_address",
				table:  "requests",
			},
			wantQuery: `SELECT DISTINCT id, user_id, first_name, last_name, address_main, address_secondary, city, zip_code, state, country, phone, email, comment, status, start_at, stop_at, created_at, updated_at, account_address FROM requests LIMIT 100`,
		},
		{
			name: "select all with Limit, Offset, Sort",
			filter: NewFilter().
				SetLimit(100).
				SetOffset(10).
				SetSort("id", DirectionDescending),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests ORDER BY id DESC LIMIT 100 OFFSET 10",
		},
		{
			name: "select all with Limit, Offset, Sort, GroupBy",
			filter: NewFilter().
				SetLimit(100).
				SetOffset(10).
				SetSort("id", DirectionDescending).
				SetGroupBy("address_main"),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests GROUP BY address_main ORDER BY id DESC LIMIT 100 OFFSET 10",
		},
		{
			name: "select with one predicator",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateLt{"amount": 20}),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND amount < $2",
			wantArgs:  []any{"active", 20},
		},
		{
			name: "select with multiple predicates",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateLt{"amount": 20}, PredicateGtOrEq{"price": 40}),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND amount < $2 AND price >= $3",
			wantArgs:  []any{"active", 20, 40},
		},
		{
			name: "select with one predicate with multiple fields",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateLt{"amount": 20, "price": 40}),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND amount < $2 AND price < $3",
			wantArgs:  []any{"active", 20, 40},
		},
		{
			name: "select where field is null",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateIsNull("level")),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND level IS NULL",
			wantArgs:  []any{"active"},
		},
		{
			name: "select where multiple fields are null",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateIsNull("level", "name")),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND level IS NULL AND name IS NULL",
			wantArgs:  []any{"active"},
		},
		{
			name: "select where one field is null and one is not null",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateIsNull("level", "name"), PredicateIsNotNull("surname")),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND level IS NULL AND name IS NULL AND surname IS NOT NULL",
			wantArgs:  []any{"active"},
		},
		{
			name: "select with between",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateIsNull("level", "name"), PredicateIsNotNull("surname"), PredicateBetween("age", 10, 30)),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND level IS NULL AND name IS NULL AND surname IS NOT NULL AND age BETWEEN $2 AND $3",
			wantArgs:  []any{"active", 10, 30},
		},
		{
			name: "select with not between",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateIsNull("level", "name"), PredicateNotBetween("age", 10, 30), PredicateIsNotNull("surname")),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1)) AND level IS NULL AND name IS NULL AND age NOT BETWEEN $2 AND $3 AND surname IS NOT NULL",
			wantArgs:  []any{"active", 10, 30},
		},
		{
			name: "select where field is not in",
			filter: NewFilter().
				SetCondition(ConditionAND).
				SetArgument("status", "active").
				SetPredicates(PredicateNotIn("status", "blocked", "deleted")).
				SetArgument("amount", 10),
			args: args{
				table: "requests",
			},
			wantQuery: "SELECT * FROM requests WHERE (status IN ($1) AND amount IN ($2)) AND status NOT IN ($3, $4)",
			wantArgs:  []any{"active", 10, "blocked", "deleted"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotQuery, gotArgs := tt.filter.Build(tt.args.table, tt.args.fields)
			if tt.args.fields == "" {
				gotQuery, gotArgs = tt.filter.Build(tt.args.table)
			}

			if !reflect.DeepEqual(gotQuery, tt.wantQuery) {
				t.Errorf("\n Query got = %v \n want      = %v", gotQuery, tt.wantQuery)
			}

			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("\n Arguments got = %v,\n want          = %v", gotArgs, tt.wantArgs)
			}
		})
	}
}

func Test_Exclude(t *testing.T) {
	params := NewFilter().SetArgumentsMap(map[string][]any{"status": {"active"}, "foo": {"bar"}, "user_id": {"1"}})

	userReq := params.ExcludeArguments("status")

	t.Log(params.GetArguments("status")...)

	t.Log(userReq.GetArguments("status")...)
}

// BenchmarkExclude-12    	 5471670	       202.0 ns/op
func BenchmarkExclude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		params := NewFilter().SetArgumentsMap(map[string][]any{"status": {"active"}, "foo": {"bar"}, "user_id": {"1"}})
		_ = params.ExcludeArguments("status", "user_id", "non exists")
	}
}
