/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by  Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/callisto/tree/dev/LICENSES
 */

package database

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"

	"github.com/stalwart-algoritmiclab/callisto/database/stwart/chain/secured"
)

const (
	host       = "localhost"
	dbName     = "callisto"
	dbUser     = "postgres"
	dbPassword = "postgres"
)

const (
	TestAddressCreator = "stwart1y3n5h0r5nwvw6n6dfk0xn3xxh4thph9nda6t98"
)

var (
	DB    *sqlx.DB
	Codec codec.Codec

	Datastore struct {
		Secured *secured.Repository
	}
)

func init() {
	var err error

	DB, err = sqlx.Connect("pgx", fmt.Sprintf("host=%s port=5432 user=%s password=%s dbname=%s sslmode=disable",
		host,
		dbUser,
		dbPassword,
		dbName,
	))
	if err != nil {
		log.Fatal().Err(err)
		panic(err)
	}

	// Create the codec.
	// TODO: rework it: Codec = registrar.Context{}.EncodingConfig.Codec

	// Custom modules
	Datastore.Secured = secured.NewRepository(DB, Codec)
}
