SET statement_timeout = 0;

--bun:split

create table courses
(
    id                uuid primary key not null,

    name              varchar(45)      not null,
    slug              varchar(60)      not null,
    short_description varchar(400)     not null,
    description       varchar(3000)    not null,
    logo              varchar(800)     not null,
    banner            varchar(800)     not null,

    updated_at        timestamptz,
    deleted_at        timestamptz,
    created_at        timestamptz      not null
);

create unique index courses_slug_idx on courses (slug) where deleted_at is null;

--bun:split

create table links
(
    id         uuid primary key not null,

    name       varchar(15)      not null,
    link       varchar(800)     not null,
    course_id  uuid             not null,

    updated_at timestamptz,
    deleted_at timestamptz,
    created_at timestamptz      not null,

    foreign key (course_id) references courses (id)
);

create index links_course_id_idx on links (course_id);

--bun:split

create table course_author
(
    course_id uuid not null,
    user_id   uuid not null,

    foreign key (course_id) references courses (id),
    foreign key (user_id) references users (id),

    primary key (course_id, user_id)
);

--bun:split

create table course_student
(
    course_id uuid not null,
    user_id   uuid not null,

    foreign key (course_id) references courses (id),
    foreign key (user_id) references users (id),
    primary key (course_id, user_id)
);