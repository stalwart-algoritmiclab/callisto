-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_referrals_create_user
(
    id              BIGINT PRIMARY KEY,
    creator         TEXT   NOT NULL,
    account_address TEXT   NOT NULL,
    referrer        TEXT   NOT NULL,
    referrals       TEXT[] NOT NULL,
    tx_hash         TEXT   NOT NULL
);

CREATE INDEX referrals_create_user_creator_index ON stwart_referrals_create_user (creator);
CREATE INDEX referrals_create_user_tx_hash_index ON stwart_referrals_create_user (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_referrals_update_user
(
    id              BIGINT PRIMARY KEY,
    creator         TEXT   NOT NULL,
    account_address TEXT   NOT NULL,
    referrer        TEXT   NOT NULL,
    referrals       TEXT[] NOT NULL,
    tx_hash         TEXT   NOT NULL
);

CREATE INDEX referrals_update_user_creator_index ON stwart_referrals_update_user (creator);
CREATE INDEX referrals_update_user_tx_hash_index ON stwart_referrals_update_user (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_referrals_delete_user
(
    id              BIGINT PRIMARY KEY,
    creator         TEXT NOT NULL,
    account_address TEXT NOT NULL,
    tx_hash         TEXT NOT NULL
);

CREATE INDEX referrals_delete_user_creator_index ON stwart_referrals_delete_user (creator);
CREATE INDEX referrals_delete_user_tx_hash_index ON stwart_referrals_delete_user (tx_hash);

CREATE TABLE IF NOT EXISTS stwart_referrals_set_referrer
(
    id       BIGINT PRIMARY KEY,
    creator  TEXT NOT NULL,
    referrer TEXT NOT NULL,
    referral TEXT NOT NULL,
    tx_hash  TEXT NOT NULL
);

CREATE INDEX referrals_set_referrer_creator_index ON stwart_referrals_set_referrer (creator);
CREATE INDEX referrals_set_referrer_tx_hash_index ON stwart_referrals_set_referrer (tx_hash);

-- +migrate Down
DROP TABLE IF EXISTS stwart_referrals_create_user;
DROP TABLE IF EXISTS stwart_referrals_update_user;
DROP TABLE IF EXISTS stwart_referrals_delete_user;
DROP TABLE IF EXISTS stwart_referrals_set_referrer;
DROP INDEX IF EXISTS referrals_create_user_creator_index;
DROP INDEX IF EXISTS referrals_create_user_tx_hash_index;
DROP INDEX IF EXISTS referrals_update_user_creator_index;
DROP INDEX IF EXISTS referrals_update_user_tx_hash_index;
DROP INDEX IF EXISTS referrals_delete_user_creator_index;
DROP INDEX IF EXISTS referrals_delete_user_tx_hash_index;
DROP INDEX IF EXISTS referrals_set_referrer_creator_index;
DROP INDEX IF EXISTS referrals_set_referrer_tx_hash_index;
