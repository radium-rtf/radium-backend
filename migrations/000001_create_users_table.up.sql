CREATE TABLE IF NOT EXISTS users
(
    id           uuid PRIMARY KEY,
    name         VARCHAR(100)             NOT NULL,
    username     VARCHAR(50) UNIQUE       NOT NULL,
    password     VARCHAR(400)             NOT NULL,
    email        VARCHAR(300) UNIQUE      NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now()
);