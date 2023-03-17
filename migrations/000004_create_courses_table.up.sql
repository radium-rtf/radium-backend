CREATE TABLE IF NOT EXISTS courses
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(200) UNIQUE      NOT NULL,
    description VARCHAR(3000)            not null,
    logo        VARCHAR(400)             NOT NULL,
    chat        VARCHAR(400)             NOT NULL,
    type        VARCHAR(400)             NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);