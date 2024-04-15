SET statement_timeout = 0;

--bun:split

create type notification_type as enum ('review', 'deadline', 'info');

create table if not exists notifications
(
    id uuid not null primary key,
    user_id uuid not null references users(id),
    answer_id uuid not null references answers(id),
    type notification_type not null,
    read boolean not null default false,

    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz      not null
)
