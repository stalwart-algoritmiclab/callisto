-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_faucet
(
    id       SERIAL NOT NULL PRIMARY KEY,
    creator  TEXT   NOT NULL,
    tx_hash  TEXT   NOT NULL UNIQUE,
    address  TEXT   NOT NULL
);

CREATE INDEX faucet_creator_index ON stwart_faucet (creator);
CREATE INDEX faucet_tx_hash_index ON stwart_faucet (tx_hash);

-- +migrate Down
DROP INDEX IF EXISTS faucet_tx_hash_index;
DROP INDEX IF EXISTS faucet_creator_index;
DROP TABLE IF EXISTS stwart_faucet;