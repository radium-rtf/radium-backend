SET statement_timeout = 0;

--bun:split

alter table courses alter column name type varchar(128);
alter table courses alter column slug type varchar(700);