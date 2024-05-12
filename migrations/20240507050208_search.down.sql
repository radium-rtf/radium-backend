drop index if exists idx_rum_users_email;
drop index if exists idx_rum_users_name;

drop index if exists idx_rum_groups_name;
drop index if exists idx_rum_courses_name;

drop trigger if exists name_tsvector_update_trigger on courses;
drop trigger if exists name_tsvector_update_trigger on groups;

drop trigger if exists name_tsvector_update_trigger on users;
drop trigger if exists email_tsvector_update_trigger on users;

alter table courses drop column if exists tsvector_name;
alter table groups drop column if exists tsvector_name;

alter table users drop column if exists tsvector_name;
alter table users drop column if exists tsvector_email;

drop extension rum;