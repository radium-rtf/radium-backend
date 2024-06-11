SET statement_timeout = 0;

--bun:split

CREATE TABLE wave.groups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT,
    avatar_url TEXT,
    owner_id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
    -- FOREIGN KEY (owner_id) REFERENCES radium.users (id),
);

--bun:split

CREATE TABLE wave.group_member (
    group_id UUID,
    user_id UUID,
    admin BOOLEAN NOT NULL,
    PRIMARY KEY (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES wave.groups (id)
    -- FOREIGN KEY (user_id) REFERENCES radium.users (id)
);

--bun:split

CREATE TABLE wave.group_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);
