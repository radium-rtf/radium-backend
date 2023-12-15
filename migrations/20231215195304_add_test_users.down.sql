SET statement_timeout = 0;

--bun:split

delete from sessions
where user_id = '11af02da-bf9e-4769-aa07-36903517733c';

delete from roles
where user_id = '11af02da-bf9e-4769-aa07-36903517733c';

delete from group_student
where user_id = '11af02da-bf9e-4769-aa07-36903517733c';

delete from users
where id = '11af02da-bf9e-4769-aa07-36903517733c';
