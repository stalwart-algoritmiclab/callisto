-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_rates_create_addresses
(
    id      SERIAL NOT NULL PRIMARY KEY,
    creator TEXT   NOT NULL,
    address TEXT[] NOT NULL,
    tx_hash TEXT   NOT NULL
);

CREATE INDEX rates_create_addresses_creator_index ON stwart_rates_create_addresses (creator);
CREATE INDEX rates_create_addresses_tx_hash_index ON stwart_rates_create_addresses (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_rates_update_addresses
(
    id         SERIAL NOT NULL PRIMARY KEY,
    creator    TEXT   NOT NULL,
    address_id BIGINT NOT NULL,
    address    TEXT[] NOT NULL,
    tx_hash    TEXT   NOT NULL
);

CREATE INDEX rates_update_addresses_creator_index ON stwart_rates_update_addresses (creator);
CREATE INDEX rates_update_addresses_tx_hash_index ON stwart_rates_update_addresses (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_rates_delete_addresses
(
    id         SERIAL NOT NULL PRIMARY KEY,
    creator    TEXT   NOT NULL,
    address_id BIGINT NOT NULL,
    tx_hash    TEXT   NOT NULL
);

CREATE INDEX rates_delete_addresses_creator_index ON stwart_rates_delete_addresses (creator);
CREATE INDEX rates_delete_addresses_tx_hash_index ON stwart_rates_delete_addresses (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_rates_create_rates
(
    id       SERIAL  NOT NULL PRIMARY KEY,
    creator  TEXT    NOT NULL,
    decimals NUMERIC NOT NULL,
    denom    TEXT    NOT NULL,
    rate     TEXT    NOT NULL,
    tx_hash  TEXT    NOT NULL
);

CREATE INDEX rates_create_rates_creator_index ON stwart_rates_create_rates (creator);
CREATE INDEX rates_create_rates_tx_hash_index ON stwart_rates_create_rates (tx_hash);


CREATE TABLE IF NOT EXISTS stwart_rates_update_rates
(
    id      SERIAL NOT NULL PRIMARY KEY,
    creator TEXT   NOT NULL,
    denom   TEXT   NOT NULL,
    rate    TEXT   NOT NULL,
    tx_hash TEXT   NOT NULL
);

CREATE INDEX rates_update_rates_creator_index ON stwart_rates_update_rates (creator);
CREATE INDEX rates_update_rates_tx_hash_index ON stwart_rates_update_rates (tx_hash);


CREATE TABLE IF NOT EXISTS stwart_rates_delete_rates
(
    id      SERIAL NOT NULL PRIMARY KEY,
    creator TEXT   NOT NULL,
    denom   TEXT   NOT NULL,
    tx_hash TEXT   NOT NULL
);

CREATE INDEX rates_delete_rates_creator_index ON stwart_rates_delete_rates (creator);
CREATE INDEX rates_delete_rates_tx_hash_index ON stwart_rates_delete_rates (tx_hash);

-- +migrate Down
DROP INDEX rates_delete_rates_tx_hash_index;
DROP INDEX rates_delete_rates_creator_index;
DROP TABLE IF EXISTS stwart_rates_delete_rates;

DROP INDEX rates_create_rates_tx_hash_index;
DROP INDEX rates_create_rates_creator_index;
DROP TABLE IF EXISTS stwart_rates_create_rates;

DROP INDEX rates_create_addresses_tx_hash_index;
DROP INDEX rates_create_addresses_creator_index;
DROP TABLE IF EXISTS stwart_rates_create_addresses;

DROP INDEX rates_update_rates_tx_hash_index;
DROP INDEX rates_update_rates_creator_index;
DROP TABLE IF EXISTS stwart_rates_update_rates;

DROP INDEX rates_update_addresses_tx_hash_index;
DROP INDEX rates_update_addresses_creator_index;
DROP TABLE IF EXISTS stwart_rates_update_addresses;

DROP INDEX rates_delete_addresses_tx_hash_index;
DROP INDEX rates_delete_addresses_creator_index;
DROP TABLE IF EXISTS stwart_rates_delete_addresses;