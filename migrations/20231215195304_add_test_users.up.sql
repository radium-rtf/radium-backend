SET statement_timeout = 0;

--bun:split

insert into users (id, email, avatar, name, password,
                   updated_at, deleted_at, created_at) values
            ('11af02da-bf9e-4769-aa07-36903517733c', 'test.test.test@urfu.me', '', 'test', '$2a$10$DtNkeXx9KGmq2CdpO8USKuQTonOEX2ClVyBqq9NhlG7eT1Xi0EwJS',
             null, null, now()) on conflict do nothing;

insert into roles (user_id, is_teacher, is_author, is_coauthor) values
                ('11af02da-bf9e-4769-aa07-36903517733c', true, true, true) on conflict do nothing;

insert into group_student (group_id, user_id) values
                ('81af02da-bf9e-4769-aa07-36903517733d', '11af02da-bf9e-4769-aa07-36903517733c') on conflict do nothing;