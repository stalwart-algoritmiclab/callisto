-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_core_fees
(
    id        SERIAL NOT NULL PRIMARY KEY,
    tx_hash   TEXT   NOT NULL UNIQUE,
    creator   TEXT   NOT NULL,
    comission TEXT   NOT NULL,
    address   TEXT   NOT NULL
);

CREATE INDEX core_fees_creator_index ON stwart_core_fees (creator);
CREATE INDEX core_fees_tx_hash_index ON stwart_core_fees (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_core_issue
(
    id      SERIAL NOT NULL PRIMARY KEY,
    tx_hash TEXT   NOT NULL UNIQUE,
    creator TEXT   NOT NULL,
    denom   TEXT   NOT NULL,
    amount  TEXT   NOT NULL,
    address TEXT   NOT NULL
);

CREATE INDEX core_issue_creator_index ON stwart_core_issue (creator);
CREATE INDEX core_issue_tx_hash_index ON stwart_core_issue (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_core_refreward
(
    id       SERIAL NOT NULL PRIMARY KEY,
    tx_hash  TEXT   NOT NULL UNIQUE,
    creator  TEXT   NOT NULL,
    amount   TEXT   NOT NULL,
    referrer TEXT   NOT NULL
);

CREATE INDEX core_refreward_creator_index ON stwart_core_refreward (creator);
CREATE INDEX core_refreward_tx_hash_index ON stwart_core_refreward (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_core_refund
(
    id           SERIAL NOT NULL PRIMARY KEY,
    tx_hash      TEXT   NOT NULL UNIQUE,
    creator      TEXT   NOT NULL,
    amount       TEXT   NOT NULL,
    from_address TEXT   NOT NULL,
    to_address   TEXT   NOT NULL
);

CREATE INDEX core_refund_creator_index ON stwart_core_refund (creator);
CREATE INDEX core_refund_tx_hash_index ON stwart_core_refund (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_core_send
(
    id      SERIAL NOT NULL PRIMARY KEY,
    tx_hash TEXT   NOT NULL UNIQUE,
    creator TEXT   NOT NULL,
    from_address TEXT   NOT NULL,
    to_address   TEXT   NOT NULL,
    amount       TEXT   NOT NULL,
    denom        TEXT   NOT NULL
);

CREATE INDEX core_send_creator_index ON stwart_core_send (creator);
CREATE INDEX core_send_tx_hash_index ON stwart_core_send (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_core_withdraw
(
    id      SERIAL NOT NULL PRIMARY KEY,
    tx_hash TEXT   NOT NULL UNIQUE,
    creator TEXT   NOT NULL,
    amount  TEXT   NOT NULL,
    denom   TEXT   NOT NULL,
    address TEXT   NOT NULL
);

CREATE INDEX core_withdraw_creator_index ON stwart_core_withdraw (creator);
CREATE INDEX core_withdraw_tx_hash_index ON stwart_core_withdraw (tx_hash);

-- +migrate Down

DROP INDEX IF EXISTS core_fees_tx_hash_index;
DROP INDEX IF EXISTS core_fees_creator_index;
DROP TABLE IF EXISTS stwart_core_fees;

DROP INDEX IF EXISTS core_issue_tx_hash_index;
DROP INDEX IF EXISTS core_issue_creator_index;
DROP TABLE IF EXISTS stwart_core_issue;

DROP INDEX IF EXISTS core_refreward_tx_hash_index;
DROP INDEX IF EXISTS core_refreward_creator_index;
DROP TABLE IF EXISTS stwart_core_refreward;

DROP INDEX IF EXISTS core_refund_tx_hash_index;
DROP INDEX IF EXISTS core_refund_creator_index;
DROP TABLE IF EXISTS stwart_core_refund;

DROP INDEX IF EXISTS core_send_tx_hash_index;
DROP INDEX IF EXISTS core_send_creator_index;
DROP TABLE IF EXISTS stwart_core_send;

DROP INDEX IF EXISTS core_withdraw_tx_hash_index;
DROP INDEX IF EXISTS core_withdraw_creator_index;
DROP TABLE IF EXISTS stwart_core_withdraw;
