SET statement_timeout = 0;

--bun:split

update courses
set slug = substring(id::varchar, length(id::text) - 10)
where true;

update modules
set slug = substring(id::varchar, length(id::text) - 10)
where true;

update pages
set slug = substring(id::varchar, length(id::text) - 10)
where true;

create unique index if not exists courses_slug_idx on courses (slug);
create unique index if not exists pages_slug_idx on pages (slug);
create unique index if not exists modules_slug_idx on modules (slug);
