DROP TABLE IF EXISTS course_links;

ALTER TABLE IF EXISTS course
ADD COLUMN IF NOT EXISTS chat varchar(250) not null default '';