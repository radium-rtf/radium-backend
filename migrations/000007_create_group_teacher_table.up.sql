CREATE TABLE IF NOT EXISTS group_teacher
(
    id UUID PRIMARY KEY NOT NULL ,
    user_id UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL ,
    group_id UUID  REFERENCES groups (id) ON DELETE CASCADE NOT NULL ,
    UNIQUE (user_id, group_id)
)