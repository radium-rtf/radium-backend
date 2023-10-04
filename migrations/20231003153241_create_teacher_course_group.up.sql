SET statement_timeout = 0;

--bun:split

create table teacher_course_group
(
    user_id uuid not null,
    course_id uuid not null,
    group_id uuid not null,

    primary key (user_id, course_id, group_id),

    foreign key (user_id) references users(id),
    foreign key (group_id) references groups(id),
    foreign key (course_id) references courses(id)
)