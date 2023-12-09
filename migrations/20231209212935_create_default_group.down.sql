SET statement_timeout = 0;

--bun:split

delete from group_course
where group_id = '81af02da-bf9e-4769-aa07-36903517733d';

delete from group_student
where group_id = '81af02da-bf9e-4769-aa07-36903517733d';

delete from groups
where id = '81af02da-bf9e-4769-aa07-36903517733d';
