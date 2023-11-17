SET statement_timeout = 0;

--bun:split

alter table roles drop column is_coauthor;

--bun:split

drop table course_coauthor;