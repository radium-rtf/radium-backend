create table if not exists slides
(
    id serial primary key,
    name_eng varchar(30),
    name varchar(30) not null,
    module_id serial references modules(id),
    unique (name_eng, module_id)
);

create view module_slides_view as
(
select modules.id,
       modules.name,
       modules.name_eng,
       (select array_to_json(array_agg(row_to_json(slide.*))) as slides
        from (select slides.id, slides.name, slides.name_eng from slides where slides.module_id = modules.id) as slide)
from modules
    )