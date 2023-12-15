SET statement_timeout = 0;

--bun:split

insert into teacher_course_group
select group_student.user_id,
       group_course.course_id,
       group_student.group_id
from group_student
    join group_course on group_course.group_id = group_student.group_id
where group_student.group_id = '81af02da-bf9e-4769-aa07-36903517733d' on conflict do nothing;
