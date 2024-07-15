-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_exchanger
(
    id       BIGSERIAL NOT NULL PRIMARY KEY,
    creator  TEXT   NOT NULL,
    tx_hash  TEXT   NOT NULL UNIQUE,
    denom    TEXT   NOT NULL,
    amount   TEXT   NOT NULL,
    denom_to TEXT   NOT NULL
);

CREATE INDEX exchanger_creator_index ON stwart_exchanger (creator);
CREATE INDEX exchanger_tx_hash_index ON stwart_exchanger (tx_hash);

-- +migrate Down
DROP INDEX IF EXISTS exchanger_tx_hash_index;
DROP INDEX IF EXISTS exchanger_creator_index;
DROP TABLE IF EXISTS stwart_exchanger;
