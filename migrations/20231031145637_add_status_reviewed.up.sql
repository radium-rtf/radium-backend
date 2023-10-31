SET statement_timeout = 0;

--bun:split

alter type verdict_type add value if not exists 'REVIEWED'
