-- +migrate Up
CREATE TABLE top_accounts
(
    address    TEXT   NOT NULL REFERENCES account (address) PRIMARY KEY,
    type       TEXT,
    available  BIGINT          DEFAULT 0,
    delegation BIGINT          DEFAULT 0,
    unbonding  BIGINT          DEFAULT 0,
    reward     BIGINT          DEFAULT 0,
    ssc        BIGINT          DEFAULT 0,
    sum        BIGINT NOT NULL DEFAULT 0,
    height     BIGINT NOT NULL
);
CREATE INDEX top_accounts_sum_index ON top_accounts (sum);
CREATE INDEX top_accounts_height_index ON top_accounts (height);
CREATE INDEX top_accounts_type_index ON top_accounts (type);
CREATE INDEX top_accounts_ssc_index ON top_accounts (ssc);


CREATE TABLE top_accounts_params
(
    one_row_id     BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    total_accounts BIGINT  NOT NULL,
    height         BIGINT  NOT NULL,
    CHECK (one_row_id)
);
CREATE INDEX top_accounts_params_height_index ON top_accounts_params (height);

-- +migrate Down
DROP INDEX IF EXISTS top_accounts_sum_index;
DROP INDEX IF EXISTS top_accounts_height_index;
DROP INDEX IF EXISTS top_accounts_type_index;
DROP INDEX IF EXISTS top_accounts_ssc_index;
DROP TABLE IF EXISTS top_accounts;
DROP INDEX IF EXISTS top_accounts_params_height_index;
DROP TABLE IF EXISTS top_accounts_params;