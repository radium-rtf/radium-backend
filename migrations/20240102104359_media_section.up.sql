SET statement_timeout = 0;

--bun:split


alter type section_type add value 'media';

alter table sections add column url varchar(1000) references files(url) default null;

