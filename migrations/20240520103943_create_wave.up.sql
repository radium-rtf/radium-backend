SET statement_timeout = 0;

--bun:split

CREATE SCHEMA IF NOT EXISTS wave;

--bun:split

CREATE TABLE wave.contents (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    file_id UUID,
    text TEXT,
    type TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
    -- FOREIGN KEY (file_id) REFERENCES radium.files (id)
);

--bun:split

CREATE TABLE wave.messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sender_id UUID NOT NULL,
    content_id UUID NOT NULL,
    parent_message_id UUID,
    type TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (content_id) REFERENCES wave.contents (id)
    -- FOREIGN KEY (parent_message_id) REFERENCES wave.messages (id)
);

--bun:split

CREATE TABLE wave.read_by (
    message_id UUID NOT NULL,
    user_id UUID NOT NULL,
    PRIMARY KEY (message_id, user_id),
    FOREIGN KEY (message_id) REFERENCES wave.messages (id)
    -- FOREIGN KEY (user_id) REFERENCES radium.users (id)
);

--bun:split

CREATE TABLE wave.reactions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    message_id UUID NOT NULL,
    user_id UUID NOT NULL,
    reaction TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    FOREIGN KEY (message_id) REFERENCES wave.messages (id)
    -- FOREIGN KEY (user_id) REFERENCES radium.users (id)
);

--bun:split

CREATE TABLE wave.dialogues (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_user_id UUID NOT NULL,
    second_user_id UUID NOT NULL,
    UNIQUE(first_user_id, second_user_id)
);

--bun:split

CREATE TABLE wave.dialogue_settings (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ
);

--bun:split

CREATE TABLE wave.dialogue_message (
    dialogue_id UUID NOT NULL,
    message_id UUID NOT NULL,
    PRIMARY KEY (dialogue_id, message_id),
    FOREIGN KEY (dialogue_id) REFERENCES wave.dialogues (id),
    FOREIGN KEY (message_id) REFERENCES wave.messages (id)
);

--bun:split

CREATE TABLE wave.dialogue_pinned (
    dialogue_id UUID NOT NULL,
    message_id UUID NOT NULL,
    PRIMARY KEY (dialogue_id, message_id),
    FOREIGN KEY (dialogue_id) REFERENCES wave.dialogues (id),
    FOREIGN KEY (message_id) REFERENCES wave.messages (id)
);
