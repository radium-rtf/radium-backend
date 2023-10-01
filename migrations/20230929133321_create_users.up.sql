SET statement_timeout = 0;

--bun:split

create table users
(
    id         uuid primary key default gen_random_uuid(),

    email      varchar(200) not null,
    avatar     varchar(500),
    name       varchar(50)  not null,
    password   varchar(600) not null,

    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz  not null
);

create unique index users_email_idx on users (email) where deleted_at is null;

--bun:split

create table roles
(
    user_id    uuid primary key,

    is_teacher boolean not null default false,
    is_author  boolean not null default false,

    foreign key(user_id) references users(id)
);

--bun:split

create table sessions
(
    refresh_token uuid primary key not null,
    expires_in    timestamptz      not null,
    user_id       uuid             not null,
    foreign key (user_id) references users (id)
);

create index sessions_user_id_idx on sessions (user_id);
