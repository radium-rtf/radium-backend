SET statement_timeout = 0;

--bun:split

alter table reviews add column comment varchar(500) default '';
