SET statement_timeout = 0;

--bun:split

alter table answers drop if exists file_url;

--bun:split

alter table sections drop if exists file_types;

--bun:split

drop table if exists files;
