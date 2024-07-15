-- +migrate Up
CREATE TABLE IF NOT EXISTS stwart_referrals_set_referrer
(
    id       BIGSERIAL NOT NULL PRIMARY KEY,
    creator  TEXT NOT NULL,
    referrer TEXT NOT NULL,
    referral TEXT NOT NULL,
    tx_hash  TEXT NOT NULL UNIQUE
);

CREATE INDEX referrals_set_referrer_creator_index ON stwart_referrals_set_referrer (creator);
CREATE INDEX referrals_set_referrer_tx_hash_index ON stwart_referrals_set_referrer (tx_hash);

-- +migrate Down
DROP TABLE IF EXISTS stwart_referrals_set_referrer;
DROP INDEX IF EXISTS referrals_set_referrer_creator_index;
DROP INDEX IF EXISTS referrals_set_referrer_tx_hash_index;
