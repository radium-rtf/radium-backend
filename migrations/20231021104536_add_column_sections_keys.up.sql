SET statement_timeout = 0;

--bun:split

alter table sections
    add keys varchar(150)[] not null default array[]::varchar[];

--bun:split

alter type section_type add value 'mapping'