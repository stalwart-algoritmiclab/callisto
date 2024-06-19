-- +migrate Up
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

CREATE TABLE stwart_feepolicy_msg_create_tariffs
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    denom   TEXT      NOT NULL,
    creator TEXT      NOT NULL,
    tx_hash TEXT      NOT NULL UNIQUE,
    height  BIGINT    NOT NULL,
    FOREIGN KEY (tx_hash) REFERENCES stwart_feepolicy_tariffs (tx_hash) ON DELETE CASCADE
);

CREATE TABLE stwart_feepolicy_msg_update_tariffs
(
    id      BIGSERIAL NOT NULL PRIMARY KEY,
    creator TEXT      NOT NULL,
    denom   TEXT      NOT NULL,
    tx_hash TEXT      NOT NULL UNIQUE,
    height  BIGINT    NOT NULL,
    FOREIGN KEY (tx_hash) REFERENCES stwart_feepolicy_tariffs (tx_hash) ON DELETE CASCADE
);

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
-- +migrate Down
DROP TABLE IF EXISTS stwart_feepolicy_msg_delete_tariffs;
DROP TABLE IF EXISTS stwart_feepolicy_msg_update_tariffs;
DROP TABLE IF EXISTS stwart_feepolicy_msg_create_tariffs;
DROP TABLE IF EXISTS stwart_feepolicy_tariffs;
