SET statement_timeout = 0;

--bun:split

alter table teacher_course_group rename to teacher;

create table if not exists students
(
    user_id uuid not null references users(id),
    course_id uuid not null references courses(id),
    group_id uuid not null references groups(id),
    primary key (user_id, course_id)
);

insert into students (user_id, course_id, group_id)
select c.user_id, c.course_id, '81af02da-bf9e-4769-aa07-36903517733d' from course_student as c
on conflict do nothing;

drop table if exists course_student;

drop table if exists group_student;