-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_secured_addresses
(
    id        SERIAL   NOT NULL PRIMARY KEY,
    creator   TEXT     NOT NULL,
    tx_hash   TEXT     NOT NULL,
    addresses TEXT[]   NOT NULL
);

CREATE INDEX secured_creator_index ON stwart_secured_addresses (creator);
CREATE INDEX secured_tx_hash_index ON stwart_secured_addresses (tx_hash);

-- +migrate Down
DROP INDEX IF EXISTS secured_tx_hash_index;
DROP INDEX IF EXISTS secured_creator_index;
DROP TABLE IF EXISTS stwart_secured_addresses;
