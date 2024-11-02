-- +goose Up
create table meta
(
    key   TEXT PRIMARY KEY,
    value TEXT
);

create table currency
(
    code          CHAR(5) PRIMARY KEY,
    symbol        VARCHAR(10) NOT NULL,
    exchange_rate NUMERIC     NOT NULL,
    last_updated  TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


create table market
(
    code          varchar(28) PRIMARY KEY,
    name          varchar(255)                                         NOT NULL,
    currency_code CHAR(5) REFERENCES currency (code) ON DELETE CASCADE NOT NULL
);

create table workspace
(
    id            BIGSERIAL PRIMARY KEY,
    name          varchar(255)                                         NOT NULL,
    currency_code CHAR(5) REFERENCES currency (code) ON DELETE CASCADE NOT NULL,
    icon          TEXT                                                 NOT NULL,
    created_at    TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

create type security_type as ENUM ('stock', 'mutual_fund', 'bond', 'etf', 'other');

create table security
(
    id         BIGSERIAL PRIMARY KEY,
    market_id  varchar(28) REFERENCES market (code) ON DELETE CASCADE NOT NULL,
    identifier VARCHAR(100)                                           NOT NULL,
    value      NUMERIC                                                NOT NULL,
    type       security_type                                          NOT NULL
);


create table "user"
(
    id       BIGSERIAL PRIMARY KEY,
    email    varchar(255) UNIQUE NOT NULL,
    icon     TEXT                NOT NULL,
    password TEXT,
    name     varchar(255) UNIQUE,
    hash     TEXT                NOT NULL
);

create type role as ENUM ('owner', 'read', 'read_write', 'none');

create table workspace_access
(
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    user_id      BIGINT REFERENCES "user" (id) ON DELETE CASCADE    NOT NULL,
    role         role                                               NOT NULL,
    PRIMARY KEY (workspace_id, user_id)
);


create table file_upload
(
    id           BIGSERIAL PRIMARY KEY,
    path         TEXT                                               NOT NULL,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    name         varchar(255)                                       NOT NULL,
    mime_type    varchar(255)                                       NOT NULL,
    created_at   timestamptz DEFAULT CURRENT_TIMESTAMP              NOT NULL,
    linked       BOOLEAN     DEFAULT false                          NOT NULL
);

create type account_type as ENUM ('cash', 'wallet', 'bank');

create table deposit_account
(
    id              BIGSERIAL PRIMARY KEY,
    workspace_id    BIGINT REFERENCES workspace (id) ON DELETE CASCADE   NOT NULL,
    user_id         BIGINT REFERENCES workspace (id) ON DELETE CASCADE   NOT NULL,
    currency_code   CHAR(5) REFERENCES currency (code) ON DELETE CASCADE NOT NULL,
    opening_balance NUMERIC   DEFAULT 0                                  NOT NULL,
    name            VARCHAR(128)                                         NOT NULL,
    icon            TEXT                                                 NOT NULL,
    type            account_type                                         NOT NULL,
    balance         NUMERIC   DEFAULT 0                                  NOT NULL,
    created_at      TIMESTAMP DEFAULT NOW()                              NOT NULL,
    json            jsonb     DEFAULT '{}'::json                         NOT NULL
);

create table credit_card_account
(
    id            BIGSERIAL PRIMARY KEY,
    user_id       BIGINT REFERENCES workspace (id) ON DELETE CASCADE   NOT NULL,
    workspace_id  BIGINT REFERENCES workspace (id) ON DELETE CASCADE   NOT NULL,
    currency_code CHAR(5) REFERENCES currency (code) ON DELETE CASCADE NOT NULL,
    balance       NUMERIC                                              NOT NULL,
    credit_limit  NUMERIC                                              NOT NULL,
    interest_rate NUMERIC(5, 2) DEFAULT 12.5                           NOT NULL
);

create table security_account
(
    id           BIGSERIAL PRIMARY KEY,
    user_id      BIGINT REFERENCES workspace (id) ON DELETE CASCADE     NOT NULL,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE     NOT NULL,
    market_id    varchar(28) REFERENCES market (code) ON DELETE CASCADE NOT NULL,
    name         VARCHAR(128)                                           NOT NULL
);

create table security_item
(
    id                  BIGSERIAL PRIMARY KEY,
    security_account_id BIGINT REFERENCES security_account (id) ON DELETE CASCADE NOT NULL,
    security_id         BIGINT REFERENCES security (id) on DELETE CASCADE         NOT NULL,
    quantity            INT                                                       NOT NULL
);


create table tag
(
    id           BIGSERIAL PRIMARY KEY,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    name         VARCHAR(100)                                       NOT NULL

);

create table category
(

    id           BIGSERIAL PRIMARY KEY,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    name         VARCHAR(100)                                       NOT NULL,
    icon         TEXT                                               NOT NULL
);

create table external_entity
(
    id           BIGSERIAL PRIMARY KEY,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    name         VARCHAR(100)                                       NOT NULL,
    icon         TEXT                                               NOT NULL
);


create table asset
(
    id             BIGSERIAL PRIMARY KEY,
    workspace_id   BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    user_id        BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    name           VARCHAR(255)                                       NOT NULL,
    icon           TEXT                                               NOT NULL,
    original_value numeric                                            NOT NULL,
    current_value  numeric                                            NOT NULL,
    acquired_at    timestamptz DEFAULT CURRENT_TIMESTAMP              NOT NULL
);

create type loan_status as ENUM ('active', 'cancelled', 'paid');

create table loan
(
    id                  BIGSERIAL PRIMARY KEY,
    workspace_id        BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    user_id             BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    type                VARCHAR(128)                                       NOT NULL,
    interest_rate       NUMERIC                                            NOT NULL,
    term_months         INT                                                NOT NULL,
    loan_amount         NUMERIC                                            NOT NULL,
    principal_amount    NUMERIC                                            NOT NULL,
    loan_balance        NUMERIC                                            NOT NULL,
    principal_balance   NUMERIC                                            NOT NULL,
    monthly_installment NUMERIC                                            NOT NULL,
    start_date          DATE                                               NOT NULL,
    status              loan_status                                        NOT NULL
);


create type transaction_source as ENUM ('deposit_account', 'credit_card', 'security_account', 'entity', 'asset', 'loan');

create table transaction
(
    id           BIGSERIAL PRIMARY KEY,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    user_id      BIGINT REFERENCES "user" (id) ON DELETE CASCADE    NOT NULL,
    amount       NUMERIC     DEFAULT 0                              NOT NULL,
    from_id      BIGINT                                             NOT NULL,
    to_id        BIGINT                                             NOT NULL,
    source       transaction_source                                 NOT NULL,
    dest         transaction_source                                 NOT NULL,
    created_at   timestamptz DEFAULT current_timestamp              NOT NULL,
    attachments  TEXT ARRAY[5],
    tag_ids      BIGINT ARRAY[10],
    category_id  BIGINT                                             REFERENCES category (id) ON DELETE SET NULL
);

create table loan_repayment
(
    id             BIGSERIAL PRIMARY KEY,
    workspace_id   BIGINT REFERENCES workspace (id) ON DELETE CASCADE   NOT NULL,
    transaction_id BIGINT REFERENCES transaction (id) ON DELETE CASCADE NOT NULL,
    user_id        BIGINT REFERENCES workspace (id) ON DELETE CASCADE   NOT NULL,
    principal_paid NUMERIC                                              NOT NULL,
    interest_paid  NUMERIC                                              NOT NULL
);


create type goal_status as ENUM ('active', 'completed', 'failed');

create table goal
(
    id           BIGSERIAL PRIMARY KEY,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    user_id      BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    amount       NUMERIC     DEFAULT 0                              NOT NULL,
    icon         TEXT                                               NOT NULL,
    target_date  timestamptz DEFAULT current_timestamp              NOT NULL,
    status       goal_status                                        NOT NULL,
    deadline     timestamptz DEFAULT current_timestamp              NOT NULL
);

create table goal_distribution
(
    id          BIGSERIAL PRIMARY KEY,
    source_id   BIGINT             NOT NULL,
    source_type transaction_source NOT NULL,
    quantity    NUMERIC            NOT NULL
);


create table envelope
(
    id           BIGSERIAL PRIMARY KEY,
    workspace_id BIGINT REFERENCES workspace (id) ON DELETE CASCADE NOT NULL,
    user_id      BIGINT                                             REFERENCES "user" (id) ON DELETE SET NULL
);

create table envelope_template
(
    id          BIGSERIAL PRIMARY KEY,
    envelope_id BIGINT REFERENCES envelope (id) ON DELETE CASCADE NOT NULL,
    category_id BIGINT REFERENCES category (id) ON DELETE CASCADE NOT NULL,
    amount      NUMERIC DEFAULT 0                                 NOT NULL,
    active      BOOLEAN DEFAULT TRUE                              NOT NULL
);

create table envelope_allocation
(
    id             BIGSERIAL PRIMARY KEY,
    month_notation INT               NOT NULL,
    template_id    BIGINT            REFERENCES envelope_template (id) ON DELETE SET NULL,
    amount         NUMERIC DEFAULT 0 NOT NULL,
    utilized       NUMERIC DEFAULT 0 NOT NULL
);

create type recurring_type as ENUM ('subscription', 'bills', 'income', 'expense');
create type transaction_frequency as ENUM ('daily', 'weekly', 'monthly', 'yearly', 'manual');

create table recurring_transaction
(
    id              BIGSERIAL PRIMARY KEY,
    category_id     BIGINT REFERENCES category (id) ON DELETE CASCADE NOT NULL,
    user_id         BIGINT REFERENCES "user" (id) ON DELETE CASCADE   NOT NULL,
    amount          NUMERIC                                           NOT NULL,
    from_id         BIGINT                                            NOT NULL,
    to_id           BIGINT                                            NOT NULL,
    source          transaction_source                                NOT NULL,
    dest            transaction_source                                NOT NULL,
    frequency       transaction_frequency                             NOT NULL,
    next_occurrence timestamptz
);


-- +goose Down
drop table recurring_transaction;
drop type transaction_frequency;
drop type recurring_type;
drop table envelope_allocation;
drop table envelope_template;
drop table envelope;
drop table goal_distribution;
drop table goal;
drop type goal_status;
drop table loan_repayment;
drop table transaction;
drop type transaction_source;
drop table loan;
drop type loan_status;
drop table asset;
drop table external_entity;
drop table category;
drop table tag;
drop table security_item;
drop table security_account;
drop table credit_card_account;
drop table deposit_account;
drop type account_type;
drop table file_upload;
drop table workspace_access;
drop type role;
drop table "user";
drop table security;
drop type security_type;
drop table workspace;
drop table market;
drop table currency;
drop table meta;