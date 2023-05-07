create table if not exists sections_question_answers
(
    id serial not null,
    section_id serial references sections_question(id) not null,
    answer text not null,
    user_id uuid references users(id) not null,
    verdict varchar(10) not null
)