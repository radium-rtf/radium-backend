SET statement_timeout = 0;

--bun:split

create type verdict_type as enum ('OK', 'WA', 'WAIT');

--bun:split

create table answers
(
    id         uuid primary key not null,

    type       section_type     not null,
    verdict    verdict_type     not null,
    language   varchar(50),

    answer     text             not null,
    answers    varchar(1000)[],

    user_id    uuid             not null,
    section_id uuid             not null,
    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz      not null,

    foreign key (section_id) references sections (id),
    foreign key (user_id) references users (id)
);

create index answers_user_section_id_idx on answers (user_id, section_id);