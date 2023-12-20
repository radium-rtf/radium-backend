SET statement_timeout = 0;

--bun:split

create unique index courses_slug_idx on courses (slug) where deleted_at is null;