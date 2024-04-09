SET statement_timeout = 0;

--bun:split

alter table courses alter column short_description type varchar(512);
alter table courses alter column description type varchar(4096);

alter table modules alter column name type varchar(48);


alter table links alter column name type varchar(64);
alter table links alter column link type varchar(1024);
