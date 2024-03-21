CREATE TABLE IF NOT EXISTS "user" (
    id            bigserial primary key,
    login         varchar(255)   not null,
    password_hash varchar(128)   not null,
    balance       numeric(12, 2) not null default 0,
    withdrawn     numeric(12, 2) not null default 0,
    CONSTRAINT login_unique UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS "order" (
    id          bigserial primary key,
    number      varchar(255)   not null,
    status      varchar(10)    not null,
    accrual     numeric(12, 2) not null default 0,
    uploaded_at Timestamp      not null,
    user_id     bigint         not null,
    CONSTRAINT number_unique UNIQUE (number)
);

CREATE TABLE IF NOT EXISTS "withdrawal" (
    id           bigserial primary key,
    "order"      varchar(255)   not null,
    sum          numeric(12, 2) not null,
    processed_at Timestamp      not null,
    user_id      bigint         not null
);
