SET statement_timeout = 0;

--bun:split

insert into groups (id, name, invite_code, created_at) values
    ('81af02da-bf9e-4769-aa07-36903517733d', 'Default Group', 'TPLETACRTSNEA', now());

insert into group_student (group_id, user_id)
select '81af02da-bf9e-4769-aa07-36903517733d', id from users;

insert into group_course (group_id, course_id)
select '81af02da-bf9e-4769-aa07-36903517733d', id from courses;