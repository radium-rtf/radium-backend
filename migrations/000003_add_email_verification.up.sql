ALTER TABLE users ADD COLUMN verification_code varchar;
ALTER TABLE users ADD COLUMN is_verified bool not null default false;