create table if not exists modules
(
    id serial primary key,
    name_eng varchar(40) not null,
    course_id serial not null references courses(id),
    name    varchar(40) not null,
    order_by bigserial not null,
    unique (name_eng, course_id)
);

create view courses_with_modules_view as
(
    select courses.id, courses.logo, courses.name, courses.type,
           (select array_to_json(array_agg(row_to_json(module.*))) as modules
            from (
                select modules.name_eng, modules.name from modules
                         where modules.course_id = courses.id
                         order by modules.order_by
                 ) as module
            )
    from courses
);