ALTER TABLE IF EXISTS courses
ADD COLUMN author_id uuid not null default '00000000-0000-0000-0000-000000000000';