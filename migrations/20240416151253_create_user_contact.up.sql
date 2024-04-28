SET statement_timeout = 0;

--bun:split

create table contacts
(
    name       varchar(800)      not null,
    link       varchar(800)     not null,
    user_id   uuid primary key not null,

    foreign key (user_id) references users (id)
);

create index contacts_user_id_idx on contacts (user_id);
