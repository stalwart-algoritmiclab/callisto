-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_secured_create_addresses
(
    id        BIGSERIAL NOT NULL PRIMARY KEY,
    creator   TEXT      NOT NULL,
    tx_hash   TEXT      NOT NULL,
    addresses TEXT[]    NOT NULL
);

CREATE INDEX secured_create_addresses_creator_index ON stwart_secured_create_addresses (creator);
CREATE INDEX secured_create_addresses_tx_hash_index ON stwart_secured_create_addresses (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_secured_update_addresses
(
    id         BIGSERIAL NOT NULL PRIMARY KEY,
    address_id BIGINT    NOT NULL,
    creator    TEXT      NOT NULL,
    tx_hash    TEXT      NOT NULL,
    addresses  TEXT[]    NOT NULL
);

CREATE INDEX secured_update_addresses_creator_index ON stwart_secured_update_addresses (creator);
CREATE INDEX secured_update_addresses_tx_hash_index ON stwart_secured_update_addresses (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_secured_delete_addresses
(
    id         BIGSERIAL  NOT NULL PRIMARY KEY,
    address_id BIGINT     NOT NULL,
    creator    TEXT       NOT NULL,
    tx_hash    TEXT       NOT NULL
);

CREATE INDEX secured_delete_addresses_creator_index ON stwart_secured_delete_addresses (creator);
CREATE INDEX secured_delete_addresses_tx_hash_index ON stwart_secured_delete_addresses (tx_hash);

-- +migrate Down
DROP INDEX IF EXISTS secured_delete_addresses_tx_hash_index;
DROP INDEX IF EXISTS secured_delete_addresses_tx_hash_index;
DROP TABLE IF EXISTS stwart_secured_delete_addresses;

DROP INDEX IF EXISTS secured_update_addresses_tx_hash_index;
DROP INDEX IF EXISTS secured_update_addresses_tx_hash_index;
DROP TABLE IF EXISTS stwart_secured_update_addresses;

DROP INDEX IF EXISTS secured_create_addresses_tx_hash_index;
DROP INDEX IF EXISTS secured_create_addresses_creator_index;
DROP TABLE IF EXISTS stwart_secured_create_addresses;
