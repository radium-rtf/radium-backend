SET statement_timeout = 0;

--bun:split

alter table teacher rename to teacher_course_group;

--bun:split

create table course_student
(
    user_id uuid not null references users(id),
    course_id uuid not null references courses(id),
    primary key (user_id, course_id)
);

insert into course_student (user_id, course_id)
select user_id, course_id from students;

drop table student;