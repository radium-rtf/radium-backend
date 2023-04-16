create table if not exists course_collaborators
(
    id uuid not null primary key,
    user_email varchar(300) not null references users(email),
    course_id serial not null references courses(id),
    unique (user_email, course_id)
);
