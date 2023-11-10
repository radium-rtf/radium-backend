SET statement_timeout = 0;

--bun:split

alter table sections add column max_attempts int2;

--bun:split

update sections
set max_attempts = 5
where sections.type != 'text';
