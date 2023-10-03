SET statement_timeout = 0;

--bun:split

create type section_type as enum ('choice', 'multiChoice', 'text', 'shortAnswer', 'answer', 'code', 'permutation');

--bun:split

create table sections
(
    id         uuid primary key not null,

    type       section_type     not null,
    content    text             not null,
    max_score  smallint         not null,
    "order"    real             not null,
    variants   varchar(500)[],

    answer     varchar(5000)     not null,
    answers    varchar(500)[],

    page_id    uuid             not null,

    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz      not null,

    foreign key (page_id) references pages (id)
);