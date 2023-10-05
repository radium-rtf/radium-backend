SET statement_timeout = 0;

--bun:split

create table reviews
(
    answer_id   uuid primary key not null,

    reviewer_id uuid             not null,
    score       real             not null,

    updated_at  timestamptz,
    deleted_at  timestamptz,
    created_at  timestamptz      not null,

    foreign key (answer_id) references answers (id),
    foreign key (reviewer_id) references users (id)
)
