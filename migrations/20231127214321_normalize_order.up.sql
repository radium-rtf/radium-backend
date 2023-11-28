SET statement_timeout = 0;

--bun:split

update sections
set "order" = (
    select count(*) + 1 from sections as s
    where sections.page_id = s.page_id and (
                sections."order" > s."order" or (abs(sections."order" - s."order") < 1e-3 and sections.created_at < s.created_at)
    )
);

--bun:split

update pages
set "order" = (
    select count(*) + 1 from pages as p
    where pages.module_id = p.module_id and (
        pages."order" > p."order" or (abs(pages."order"- p."order") <  1e-3 and pages.created_at < p.created_at))
);

--bun:split

update modules
set "order" = (
    select count(*) + 1 from modules as m
    where modules.course_id = m.course_id and (
        modules."order" > m."order" or (abs(modules."order" - m."order") < 1e-3 and modules.created_at < m.created_at))
);