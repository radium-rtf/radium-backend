SET statement_timeout = 0;

--bun:split

create table groups
(
    id uuid primary key not null,

    name varchar(40) not null,
    invite_code varchar(40) not null,

    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz      not null
);

create unique index groups_invite_code_idx on groups(invite_code) where deleted_at is null;

--bun:split

create table group_course
(
    group_id uuid not null,
    course_id uuid not null,

    primary key (course_id, group_id),

    foreign key (course_id) references courses(id),
    foreign key (group_id) references groups(id)
);

--bun:split

create table group_student
(
    group_id uuid not null,
    user_id  uuid not null,

    primary key (user_id, group_id),

    foreign key (user_id) references users (id),
    foreign key (group_id) references groups (id)
)



