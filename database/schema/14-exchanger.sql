-- +migrate Up
CREATE TABLE IF NOT EXISTS exchanger
(
    id                 SERIAL      NOT NULL PRIMARY KEY,
    creator            TEXT        NOT NULL,
    denom              TEXT        NOT NULL,
    amount             TEXT        NOT NULL,
    denom_to           TEXT        NOT NULL
);

CREATE INDEX exchanger_creator_index ON exchanger (creator);

-- +migrate Down
DROP TABLE IF EXISTS exchanger;
