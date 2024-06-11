SET statement_timeout = 0;

--bun:split

DROP TABLE IF EXISTS wave.chat_message;

--bun:split

DROP TABLE IF EXISTS wave.chats;

DROP TYPE IF EXISTS wave.chat_type;

--bun:split

DROP TABLE IF EXISTS wave.read_by;

--bun:split

DROP TABLE IF EXISTS wave.reactions;

--bun:split

DROP TABLE IF EXISTS wave.messages;

DROP TYPE IF EXISTS wave.message_type;

--bun:split

DROP TABLE IF EXISTS wave.contents;

--bun:split

DROP TABLE IF EXISTS wave.dialogues;

--bun:split

DROP TABLE IF EXISTS wave.dialogue_settings;

--bun:split

DROP SCHEMA IF EXISTS wave;
