/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package filter

import (
	"reflect"
	"strings"
	"testing"
)

func TestJoiner(t *testing.T) {
	type args struct {
		mainTable string
	}
	tests := []struct {
		name      string
		joiner    Joiner
		args      args
		wantQuery []string
		wantArgs  []interface{}
	}{
		{
			name: "[success] create join",
			joiner: NewFilter().
				SetLimit(100).
				SetOffset(1).
				SetArgument("status", "active", "inactive").
				SetSortMap(map[string]SortDirection{
					"user_id": DirectionAscending,
				}).
				SetGroupBy("status").
				SetSearch("Dima%", "name", "surname").
				SetDistinct().
				ToJoiner().
				PrepareTable("users", "name", "surname", "user_id", "id").
				PrepareTable("account", "status").
				PrepareJoinStatement("INNER JOIN account on account.id = users.id"),
			args: args{
				mainTable: "users",
			},
			wantQuery: []string{
				"SELECT DISTINCT",
				"users.name,users.surname,users.user_id,users.id",
				"account.status",
				"FROM users",
				"INNER JOIN account on account.id = users.id",
				"WHERE", "account.status IN ", "AND", "LOWER(users.name::text) LIKE", "OR", "LOWER(users.surname::text) LIKE",
				"GROUP BY account.status ORDER BY users.user_id ASC LIMIT 100 OFFSET 1",
			},
			wantArgs: []interface{}{"active", "inactive", "Dima%"},
		},
		{
			name: "[success] one join",
			joiner: NewFilter().
				SetLimit(100).
				SetOffset(1).
				ToJoiner().
				PrepareTable("requests", "id", "user_id", "first_name", "last_name").
				PrepareTable("shares", "shares_amount", "status").
				PrepareJoinStatement(`INNER JOIN shares ON requests.account_address = shares.account_address`),
			args: args{
				mainTable: "requests",
			},
			wantQuery: []string{
				"SELECT ",
				"requests.id,requests.user_id,requests.first_name,requests.last_name",
				"shares.shares_amount,shares.status",
				"INNER JOIN shares ON requests.account_address = shares.account_address LIMIT 100 OFFSET 1",
			},
		},
		{
			name: "[success] two join",
			joiner: NewFilter().
				SetLimit(100).
				SetOffset(1).
				ToJoiner().
				PrepareTable("requests", "id", "user_id", "first_name", "last_name").
				PrepareTable("shares", "shares_amount", "status").
				PrepareTable("certificates", "first_name").
				PrepareJoinStatement(`LEFT OUTER JOIN shares ON requests.account_address = shares.account_address`).
				PrepareJoinStatement(`LEFT OUTER JOIN certificates ON certificates.id = requests.id`),
			args: args{
				mainTable: "requests",
			},
			wantQuery: []string{
				"SELECT ",
				"requests.id,requests.user_id,requests.first_name,requests.last_name",
				"shares.shares_amount,shares.status",
				"certificates.first_name",
				"FROM requests",
				"LEFT OUTER JOIN shares ON requests.account_address = shares.account_address",
				"LEFT OUTER JOIN certificates ON certificates.id = requests.id",
				"LIMIT 100 OFFSET 1",
			},
		},
		{
			name: "[success] join",
			joiner: NewFilter().
				SetCondition("AND").
				SetArgument("marketing_type", "basic").
				ToJoiner().
				PrepareTable(
					"packets",
					"id", "bc_code", "coins", "price", "price_currency", "bonus", "is_active", "showroom_required",
					"marketing_type",
				).
				PrepareTable(
					"translations",
					"lang", "type",
				).
				PrepareJoinStatement("INNER JOIN translations ON packets.id = translations.foreign_id::integer"),
			args: args{
				mainTable: "packets",
			},
			wantQuery: []string{
				"SELECT ",
				"packets.id,packets.bc_code,packets.coins,packets.price,packets.price_currency,packets.bonus,packets.is_active,packets.showroom_required,packets.marketing_type",
				"translations.lang,translations.type",
				" FROM packets ",
				"INNER JOIN translations ON packets.id = translations.foreign_id::integer ",
				"WHERE (packets.marketing_type IN ($1)",
			},
			wantArgs: []interface{}{"basic"},
		},
		{
			name: "[success] join",
			joiner: NewFilter().SetGroupBy("projects.id", "nft_rates.id").
				ToJoiner().
				PrepareTable("projects",
					"id", "contract_address", "internal_name", "invest_number", "name", "token", "status",
					"initial_coins", "realization_amount", "rate_rise_time", "rate_rise_percent",
				).
				PrepareTable(
					"nft_rates",
					"id AS rate_id", "rate", "currency", "description", "created_at", "updated_at",
				).
				PrepareTable("virtual_balance_nfts", "coins").
				SetFieldModifier("coins", "COALESCE(SUM(coins), 0) AS total_issued").
				PrepareJoinStatement("LEFT OUTER JOIN virtual_balance_nfts ON projects.id = virtual_balance_nfts.project_id").
				PrepareJoinStatement(`JOIN nft_rates ON projects.id = nft_rates.project_id 
		AND nft_rates.created_at = (SELECT max(created_at) FROM nft_rates WHERE nft_rates.project_id = projects.id)`),

			args: args{
				mainTable: "projects",
			},

			wantQuery: []string{
				"SELECT ",
				"projects.id,projects.contract_address,projects.internal_name,projects.invest_number,projects.name,projects.token,projects.status,projects.initial_coins,projects.realization_amount,projects.rate_rise_time,projects.rate_rise_percent",
				"nft_rates.id AS rate_id,nft_rates.rate,nft_rates.currency,nft_rates.description,nft_rates.created_at,nft_rates.updated_at",
				"COALESCE(SUM(virtual_balance_nfts.coins), 0) AS total_issued",
				"FROM projects LEFT OUTER JOIN virtual_balance_nfts ON projects.id = virtual_balance_nfts.project_id",
				"JOIN nft_rates ON projects.id = nft_rates.project_id",
				"AND nft_rates.created_at = (SELECT max(created_at)",
				"FROM nft_rates WHERE nft_rates.project_id = projects.id)",
				"GROUP BY projects.id, nft_rates.id",
			},
		},
		{
			name: "[success] simple join with between predicator",
			joiner: NewFilter().
				SetPredicates(PredicateBetween("age", 10, 30)).
				SetLimit(100).
				SetOffset(1).
				ToJoiner().
				PrepareTable("requests", "id", "user_id", "first_name", "last_name", "age").
				PrepareTable("shares", "shares_amount", "status").
				PrepareJoinStatement(`INNER JOIN shares ON requests.account_address = shares.account_address`),
			args: args{
				mainTable: "requests",
			},
			wantQuery: []string{
				"SELECT",
				"requests.id,requests.user_id,requests.first_name,requests.last_name,requests.age",
				"shares.shares_amount,shares.status",
				"INNER JOIN shares ON requests.account_address = shares.account_address",
				"requests.age BETWEEN $1 AND $2",
				"LIMIT 100 OFFSET 1",
			},
			wantArgs: []interface{}{10, 30},
		},
		{
			name: "[success] simple join with null predicator",
			joiner: NewFilter().
				SetPredicates(PredicateIsNull("age"), PredicateIsNotNull("last_name")).
				SetLimit(100).
				SetOffset(1).
				ToJoiner().
				PrepareTable("requests", "id", "user_id", "first_name", "last_name", "age").
				PrepareTable("shares", "shares_amount", "status").
				PrepareJoinStatement(`INNER JOIN shares ON requests.account_address = shares.account_address`),
			args: args{
				mainTable: "requests",
			},
			wantQuery: []string{
				"SELECT",
				"requests.id,requests.user_id,requests.first_name,requests.last_name,requests.age",
				"shares.shares_amount,shares.status",
				"INNER JOIN shares ON requests.account_address = shares.account_address",
				"requests.age IS NULL",
				"requests.last_name IS NOT NULL",
				"LIMIT 100 OFFSET 1",
			},
			wantArgs: nil,
		},
		{
			name: "[success] simple join with compare predicator",
			joiner: NewFilter().
				SetPredicates(
					PredicateLt{"age": 30},
					PredicateGt{"balance": 20},
					PredicateGtOrEq{"id": 20},
					PredicateLtOrEq{"balance": 50},
				).
				SetLimit(100).
				SetOffset(1).
				ToJoiner().
				PrepareTable("requests", "id", "user_id", "first_name", "last_name", "age", "balance").
				PrepareTable("shares", "shares_amount", "status").
				PrepareJoinStatement(`INNER JOIN shares ON requests.account_address = shares.account_address`),
			args: args{
				mainTable: "requests",
			},
			wantQuery: []string{
				"SELECT",
				"requests.id,requests.user_id,requests.first_name,requests.last_name,requests.age,requests.balance",
				"shares.shares_amount,shares.status",
				"INNER JOIN shares ON requests.account_address = shares.account_address",
				"requests.age < $1",
				"requests.balance > $2",
				"requests.id >= $3",
				"requests.balance <= $4",
			},
			wantArgs: []interface{}{30, 20, 20, 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotQuery, gotArgs := tt.joiner.Build(tt.args.mainTable)

			for _, w := range tt.wantQuery {
				if !strings.Contains(gotQuery, w) {
					t.Errorf("\n Query got = %v,\n want = %v", gotQuery, w)
				}
			}

			if !reflect.DeepEqual(gotArgs, tt.wantArgs) {
				t.Errorf("\n Arguments got = %v,\n want = %v", gotArgs, tt.wantArgs)
			}
		})
	}
}

func Test_joiner(t *testing.T) {
	query, _ := NewFilter().
		ToJoiner().
		PrepareTable(
			"projects",
			"id", "name", "description", "status", "created_at", "updated_at",
		).
		PrepareTable(
			"virtual_balance_nfts",
			"coins",
		).
		SetFieldModifier("coins", "COALEASCE(sum(coins), 0) AS coins").
		PrepareJoinStatement("LEFT OUTER JOIN virtual_balance_nfts ON projects.id = virtual_balance_nfts.project_id").
		Build("projects")

	t.Log(query)

	if !strings.Contains(query, "COALEASCE(sum(virtual_balance_nfts.coins), 0)") {
		t.Errorf("\n Query got = %v,\n want = %v", query, "COALEASCE(sum(virtual_balance_nfts.coins), 0)")
	}
}
