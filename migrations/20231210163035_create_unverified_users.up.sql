SET statement_timeout = 0;

--bun:split

create table unverified_users
(
    id                uuid primary key,

    email             varchar(200) not null,
    avatar            varchar(500),
    name              varchar(50)  not null,
    password          varchar(600) not null,

    verification_code varchar(20)  not null,
    expires_at        timestamptz  not null,
    unique (email, verification_code)
)