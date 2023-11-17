SET statement_timeout = 0;

--bun:split

alter table roles add column is_coauthor bool;

--bun:split

create table course_coauthor
(
    course_id uuid not null,
    user_id   uuid not null,

    foreign key (course_id) references courses (id),
    foreign key (user_id) references users (id),

    primary key (course_id, user_id)
);
