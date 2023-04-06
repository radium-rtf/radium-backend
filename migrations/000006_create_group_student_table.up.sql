CREATE TABLE IF NOT EXISTS group_student
(
    id SERIAL PRIMARY KEY NOT NULL ,
    user_id SERIAL REFERENCES users (id) ON DELETE CASCADE NOT NULL ,
    group_id UUID  REFERENCES groups (id) ON DELETE CASCADE NOT NULL,
    UNIQUE (user_id, group_id)
)