create table if not exists sections_text
(
    id serial primary key not null,
    order_by int not null,
    markdown text not null,
    slide_id serial references slides(id) not null
)