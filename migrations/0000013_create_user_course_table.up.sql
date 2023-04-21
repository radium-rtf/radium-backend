create table if not exists course_student
(
    id serial primary key,
    user_id uuid references users(id),
    course_id serial references courses(id),
    unique (user_id, course_id)
)