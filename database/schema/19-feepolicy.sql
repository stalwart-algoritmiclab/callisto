-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_feepolicy_create_addresses
(
    id      SERIAL NOT NULL PRIMARY KEY,
    creator TEXT   NOT NULL,
    address TEXT   NOT NULL,
    tx_hash TEXT   NOT NULL UNIQUE
);

CREATE INDEX feepolicy_create_addresses_creator_index ON stwart_feepolicy_create_addresses (creator);
CREATE INDEX feepolicy_create_addresses_tx_hash_index ON stwart_feepolicy_create_addresses (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_feepolicy_update_addresses
(
    id         SERIAL NOT NULL PRIMARY KEY,
    creator    TEXT   NOT NULL,
    address_id BIGINT NOT NULL,
    address    TEXT   NOT NULL,
    tx_hash    TEXT   NOT NULL UNIQUE
);

CREATE INDEX feepolicy_update_addresses_creator_index ON stwart_feepolicy_update_addresses (creator);
CREATE INDEX feepolicy_update_addresses_tx_hash_index ON stwart_feepolicy_update_addresses (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_feepolicy_delete_addresses
(
    id         SERIAL NOT NULL PRIMARY KEY,
    creator    TEXT   NOT NULL,
    address_id BIGINT NOT NULL,
    tx_hash    TEXT   NOT NULL UNIQUE
);

CREATE INDEX feepolicy_delete_addresses_creator_index ON stwart_feepolicy_delete_addresses (creator);
CREATE INDEX feepolicy_delete_addresses_tx_hash_index ON stwart_feepolicy_delete_addresses (tx_hash);

CREATE TABLE stwart_feepolicy_tariffs
(
    id              SERIAL PRIMARY KEY,
    tariff_id       INT   NOT NULL,
    amount          TEXT  NOT NULL,
    denom           TEXT  NOT NULL,
    min_ref_balance TEXT  NOT NULL,
    fees            JSONB NOT NULL,
    tx_hash         TEXT  NOT NULL UNIQUE
);

CREATE INDEX feepolicy_tariffs_tx_hash_index ON stwart_feepolicy_tariffs (tx_hash);

CREATE TABLE stwart_feepolicy_msg_create_tariffs
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    denom   TEXT      NOT NULL,
    creator TEXT      NOT NULL,
    tx_hash TEXT      NOT NULL UNIQUE,
    height  BIGINT    NOT NULL,
    FOREIGN KEY (tx_hash) REFERENCES stwart_feepolicy_tariffs (tx_hash) ON DELETE CASCADE
);

CREATE INDEX feepolicy_msg_create_tariffs_creator_index ON stwart_feepolicy_msg_create_tariffs (creator);
CREATE INDEX feepolicy_msg_create_tariffs_tx_hash_index ON stwart_feepolicy_msg_create_tariffs (tx_hash);

CREATE TABLE stwart_feepolicy_msg_update_tariffs
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    creator TEXT      NOT NULL,
    denom   TEXT      NOT NULL,
    tx_hash TEXT      NOT NULL UNIQUE,
    height  BIGINT    NOT NULL,
    FOREIGN KEY (tx_hash) REFERENCES stwart_feepolicy_tariffs (tx_hash) ON DELETE CASCADE
);

CREATE INDEX feepolicy_msg_update_tariffs_creator_index ON stwart_feepolicy_msg_update_tariffs (creator);
CREATE INDEX feepolicy_msg_update_tariffs_tx_hash_index ON stwart_feepolicy_msg_update_tariffs (tx_hash);

CREATE TABLE stwart_feepolicy_msg_delete_tariffs
(
    id        BIGSERIAL NOT NULL PRIMARY KEY,
    creator   TEXT      NOT NULL,
    denom     TEXT      NOT NULL,
    tariff_id TEXT      NOT NULL,
    fee_id    TEXT      NOT NULL,
    tx_hash   TEXT      NOT NULL UNIQUE,
    height    BIGINT    NOT NULL
);

CREATE INDEX feepolicy_msg_delete_tariffs_creator_index ON stwart_feepolicy_msg_delete_tariffs (creator);
CREATE INDEX feepolicy_msg_delete_tariffs_tx_hash_index ON stwart_feepolicy_msg_delete_tariffs (tx_hash);

-- +migrate Down
DROP INDEX IF EXISTS feepolicy_msg_delete_tariffs_tx_hash_index;
DROP INDEX IF EXISTS feepolicy_msg_delete_tariffs_creator_index;
DROP TABLE IF EXISTS stwart_feepolicy_msg_delete_tariffs;
DROP INDEX IF EXISTS feepolicy_msg_update_tariffs_tx_hash_index;
DROP INDEX IF EXISTS feepolicy_msg_update_tariffs_creator_index;
DROP TABLE IF EXISTS stwart_feepolicy_msg_update_tariffs;
DROP INDEX IF EXISTS feepolicy_msg_create_tariffs_tx_hash_index;
DROP INDEX IF EXISTS feepolicy_msg_create_tariffs_creator_index;
DROP TABLE IF EXISTS stwart_feepolicy_msg_create_tariffs;
DROP INDEX IF EXISTS feepolicy_tariffs_tx_hash_index;
DROP TABLE IF EXISTS stwart_feepolicy_tariffs;
DROP INDEX IF EXISTS feepolicy_delete_addresses_tx_hash_index;
DROP INDEX IF EXISTS feepolicy_delete_addresses_creator_index;
DROP TABLE IF EXISTS stwart_feepolicy_delete_addresses;
DROP INDEX IF EXISTS feepolicy_update_addresses_tx_hash_index;
DROP INDEX IF EXISTS feepolicy_update_addresses_creator_index;
DROP TABLE IF EXISTS stwart_feepolicy_update_addresses;
DROP INDEX IF EXISTS feepolicy_create_addresses_tx_hash_index;
DROP INDEX IF EXISTS feepolicy_create_addresses_creator_index;
DROP TABLE IF EXISTS stwart_feepolicy_create_addresses;