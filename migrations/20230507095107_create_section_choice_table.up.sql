create table if not exists sections_choice
(
    id serial primary key not null,
    slide_id serial references slides(id) not null,
    order_by int not null,
    cost int not null,
    question text not null,
    answer varchar(100) not null,
    variants varchar(100)[] not null
)