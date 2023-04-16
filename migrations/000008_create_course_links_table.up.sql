CREATE TABLE IF NOT EXISTS course_links
(
    id uuid PRIMARY KEY,
    course_id SERIAL REFERENCES courses (id) ON DELETE CASCADE NOT NULL,
    link varchar(250) NOT NULL,
    name varchar(100) NOT NULL,
    UNIQUE (course_id, link)
);

ALTER TABLE IF EXISTS courses
DROP COLUMN IF EXISTS chat