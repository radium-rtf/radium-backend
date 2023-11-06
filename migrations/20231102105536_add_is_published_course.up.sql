SET statement_timeout = 0;

--bun:split

alter table courses add column is_published bool default false;

--bun:split

update courses
set is_published = true;
