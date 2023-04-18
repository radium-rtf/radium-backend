create table if not exists course_modules
(
    id varchar(40) not null,
    course_id serial not null references courses(id),
    name    varchar(40) not null,
    order_by bigserial not null,
    primary key (id, course_id)
);

create view courses_with_modules_view as
(
    select courses.id, courses.logo, courses.name, courses.type,
           (select array_to_json(array_agg(row_to_json(module.*))) as modules
            from (
                select course_modules.id, course_modules.name from course_modules
                         where course_modules.course_id = courses.id
                         order by course_modules.order_by
                 ) as module
            )
    from courses
);