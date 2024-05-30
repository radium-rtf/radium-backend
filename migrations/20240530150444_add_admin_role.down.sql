SET statement_timeout = 0;

--bun:split

alter table roles drop column is_admin;
