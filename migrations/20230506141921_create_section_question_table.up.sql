create table if not exists sections_question
(
    id serial primary key not null,
    slide_id serial references slides(id) not null,
    case_sensitive bool not null,
    order_by int not null,
    cost int not null,
    question text not null,
    answer text not null
)