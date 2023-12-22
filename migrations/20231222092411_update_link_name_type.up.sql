SET statement_timeout = 0;

--bun:split

alter table links alter column name type varchar(32);
