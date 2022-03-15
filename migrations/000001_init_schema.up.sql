CREATE TABLE refresh_sessions (
    refresh_secret VARCHAR NOT NULL,
    user_id INTEGER NOT NULL,
    ip VARCHAR NOT NULL,
    user_agent VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE refresh_sessions ADD CONSTRAINT refresh_sessions_pkey PRIMARY KEY (refresh_secret);

CREATE INDEX user_idx ON refresh_sessions USING btree(user_id);