SET statement_timeout = 0;

--bun:split

create table if not exists last_visited_page
(
    user_id    uuid                     not null references users (id),
    page_id    uuid                     not null references pages (id),
    course_id  uuid                     not null references courses (id),
    updated_at timestamp with time zone not null,
    primary key (user_id, course_id)
);