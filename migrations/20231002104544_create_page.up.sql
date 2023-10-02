SET statement_timeout = 0;

--bun:split

create table pages
(
    id         uuid primary key not null,

    name       varchar(50)      not null,
    slug       varchar(60)      not null,
    "order"    real             not null,
    module_id  uuid             not null,

    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz      not null,

    foreign key (module_id) references modules (id)
);

create index pages_module_id_idx on modules(id);
