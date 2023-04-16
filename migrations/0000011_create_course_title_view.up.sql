create view courses_title_view as
(
select public.courses.id,
       public.courses.name,
       public.courses.description,
       public.courses.logo,
       public.courses.type,

       (select array_to_json(array_agg(row_to_json(link.*))) as links
        from (select public.course_links.name, public.course_links.link
              from public.course_links
              where course_links.course_id = courses.id) as link),

       (select array_to_json(array_agg(row_to_json(collaborator.*))) as collaborators
        from (select users.id, users.email, users.name, users.username
              from public.course_collaborators
                       join users on public.users.email = public.course_collaborators.user_email
              where public.course_collaborators.course_id = public.courses.id) as collaborator),

       (select row_to_json(author.*) as author
        from (select users.id, users.email, users.name, users.username
              from public.users
              where public.users.id = public.courses.author_id) as author)

from courses
    );
