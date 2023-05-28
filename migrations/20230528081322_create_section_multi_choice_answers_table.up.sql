create table if not exists sections_multi_choice_answers
(
    id serial primary key not null,
    section_id serial references sections_choice(id) not null,
    answer text[] not null,
    user_id uuid references users(id) not null,
    verdict varchar(10) not null
)