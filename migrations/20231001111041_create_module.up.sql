SET statement_timeout = 0;

--bun:split

create table modules
(
    id         uuid primary key not null,

    name       varchar(45)      not null,
    slug       varchar(60)      not null,
    "order"      real             not null,
    course_id  uuid             not null,

    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz      not null,
    foreign key (course_id) references courses (id)
);

create index modules_course_id_idx on modules(course_id);
