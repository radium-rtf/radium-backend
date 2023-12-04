SET statement_timeout = 0;

--bun:split

create table files
(
    url varchar(1000) primary key,
    name varchar(300) not null,
    type varchar(200) not null,
    size int8 not null
);

alter table answers add column file_url varchar(1000) references files(url) default null;

--bun:split

alter table sections add column file_types varchar(100)[] default null;

--bun:split

alter type section_type add value 'file';