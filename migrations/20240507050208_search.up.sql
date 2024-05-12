create extension if not exists rum;

alter table courses add column if not exists tsvector_name tsvector not null default to_tsvector('');
update courses
set tsvector_name = to_tsvector(courses.name);

CREATE TRIGGER name_tsvector_update_trigger
    BEFORE UPDATE OR INSERT ON courses
    FOR EACH ROW EXECUTE PROCEDURE tsvector_update_trigger(tsvector_name, 'pg_catalog.english', name);

CREATE INDEX if not exists idx_rum_courses_name
    ON courses
        USING rum (tsvector_name rum_tsvector_ops);



alter table groups add column if not exists tsvector_name tsvector not null default to_tsvector('');
update groups
set tsvector_name = to_tsvector(groups.name);

CREATE TRIGGER name_tsvector_update_trigger
    BEFORE UPDATE OR INSERT ON groups
    FOR EACH ROW EXECUTE PROCEDURE tsvector_update_trigger(tsvector_name, 'pg_catalog.english', name);

CREATE INDEX idx_rum_groups_name
    ON groups
        USING rum (tsvector_name rum_tsvector_ops);



alter table users add column if not exists tsvector_name tsvector not null default to_tsvector('');
update users
set tsvector_name = to_tsvector(users.name);

alter table users add column if not exists tsvector_email tsvector not null default to_tsvector('');
update users
set tsvector_email = to_tsvector(users.email);


CREATE TRIGGER email_tsvector_update_trigger
    BEFORE UPDATE OR INSERT ON users
    FOR EACH ROW EXECUTE PROCEDURE tsvector_update_trigger(tsvector_email, 'pg_catalog.english', email);

CREATE TRIGGER name_tsvector_update_trigger
    BEFORE UPDATE OR INSERT ON users
    FOR EACH ROW EXECUTE PROCEDURE tsvector_update_trigger(tsvector_name, 'pg_catalog.english', name);

CREATE INDEX if not exists idx_rum_users_name
    ON users
        USING rum (tsvector_name rum_tsvector_ops);
CREATE INDEX if not exists idx_rum_users_email
    ON users
        USING rum (tsvector_email rum_tsvector_ops);