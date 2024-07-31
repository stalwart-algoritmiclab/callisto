-- +migrate Up
CREATE TABLE stwart_polls_msg_create_polls_params
(
    id                BIGSERIAL NOT NULL PRIMARY KEY,
    creator           TEXT      NOT NULL,
    min_days_duration TEXT,
    max_days_duration TEXT,
    max_days_pending  TEXT,
    proposer_deposit  COIN[],
    burn_veto         BOOLEAN,
    tx_hash           TEXT      NOT NULL UNIQUE
);

CREATE INDEX polls_msg_create_polls_params_creator_index ON stwart_polls_msg_create_polls_params (creator);
CREATE INDEX polls_msg_create_polls_params_tx_hash_index ON stwart_polls_msg_create_polls_params (tx_hash);

CREATE TABLE stwart_polls_msg_update_polls_params
(
    id                BIGSERIAL NOT NULL PRIMARY KEY,
    creator           TEXT      NOT NULL,
    min_days_duration TEXT,
    max_days_duration TEXT,
    max_days_pending  TEXT,
    proposer_deposit  COIN[],
    burn_veto         BOOLEAN,
    tx_hash           TEXT      NOT NULL UNIQUE
);

CREATE INDEX polls_msg_update_polls_params_creator_index ON stwart_polls_msg_update_polls_params (creator);
CREATE INDEX polls_msg_update_polls_params_tx_hash_index ON stwart_polls_msg_update_polls_params (tx_hash);

CREATE TABLE stwart_polls_msg_delete_polls_params
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    creator TEXT      NOT NULL,
    tx_hash TEXT      NOT NULL UNIQUE
);

CREATE INDEX polls_msg_delete_polls_params_creator_index ON stwart_polls_msg_delete_polls_params (creator);
CREATE INDEX polls_msg_delete_polls_params_tx_hash_index ON stwart_polls_msg_delete_polls_params (tx_hash);

CREATE TABLE stwart_polls_msg_create_poll
(
    id                    BIGSERIAL NOT NULL PRIMARY KEY,
    creator               TEXT      NOT NULL,
    title                 TEXT,
    description           TEXT,
    voting_start_time     TEXT,
    voting_period         TEXT,
    min_vote_amount       BIGINT,
    min_adresses_count    BIGINT,
    min_vote_coins_amount BIGINT,
    tx_hash               TEXT      NOT NULL UNIQUE
);

CREATE INDEX polls_msg_create_poll_creator_index ON stwart_polls_msg_create_poll (creator);
CREATE INDEX polls_msg_create_poll_tx_hash_index ON stwart_polls_msg_create_poll (tx_hash);

CREATE TABLE stwart_polls_options
(
    id            BIGSERIAL NOT NULL PRIMARY KEY,
    poll_id       BIGINT    NOT NULL REFERENCES stwart_polls_msg_create_poll (id) ON DELETE CASCADE,
    voters_count  BIGINT,
    tokens_amount COIN[],
    is_veto       BOOLEAN,
    text          TEXT,
    is_vinner     BOOLEAN
);

CREATE INDEX options_poll_id_index ON stwart_polls_options (poll_id);

CREATE TABLE stwart_polls_msg_vote
(
    id        BIGSERIAL NOT NULL PRIMARY KEY,
    creator   TEXT      NOT NULL,
    poll_id   BIGINT    NOT NULL REFERENCES stwart_polls_msg_create_poll (id) ON DELETE CASCADE,
    option_id BIGINT    NOT NULL REFERENCES stwart_polls_options (id) ON DELETE CASCADE,
    amount    COIN[],
    tx_hash   TEXT      NOT NULL UNIQUE
);

CREATE INDEX vote_creator_index ON stwart_polls_msg_vote (creator);
CREATE INDEX vote_poll_id_index ON stwart_polls_msg_vote (poll_id);
CREATE INDEX vote_option_id_index ON stwart_polls_msg_vote (option_id);
CREATE INDEX vote_tx_hash_index ON stwart_polls_msg_vote (tx_hash);

-- +migrate Down
DROP INDEX IF EXISTS vote_tx_hash_index;
DROP INDEX IF EXISTS vote_option_id_index;
DROP INDEX IF EXISTS vote_poll_id_index;
DROP INDEX IF EXISTS vote_creator_index;
DROP TABLE IF EXISTS stwart_polls_msg_vote;

DROP INDEX IF EXISTS options_poll_id_index;
DROP TABLE IF EXISTS stwart_polls_options;

DROP INDEX IF EXISTS polls_msg_create_poll_tx_hash_index;
DROP INDEX IF EXISTS polls_msg_create_poll_creator_index;
DROP TABLE IF EXISTS stwart_polls_msg_create_poll;

DROP INDEX IF EXISTS polls_msg_create_polls_params_tx_hash_index;
DROP INDEX IF EXISTS polls_msg_create_polls_params_creator_index;
DROP TABLE IF EXISTS stwart_polls_msg_create_polls_params;

DROP INDEX IF EXISTS polls_msg_update_polls_params_tx_hash_index;
DROP INDEX IF EXISTS polls_msg_update_polls_params_creator_index;
DROP TABLE IF EXISTS stwart_polls_msg_update_polls_params;

DROP INDEX IF EXISTS polls_msg_delete_polls_params_tx_hash_index;
DROP INDEX IF EXISTS polls_msg_delete_polls_params_creator_index;
DROP TABLE IF EXISTS stwart_polls_msg_delete_polls_params;
