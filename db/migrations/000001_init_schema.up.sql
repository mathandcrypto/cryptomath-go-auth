CREATE TABLE "refresh_sessions" (
    "refresh_secret" VARCHAR NOT NULL PRIMARY KEY,
    "user_id" INTEGER NOT NULL,
    "ip" VARCHAR NOT NULL,
    "user_agent" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX user_idx ON "refresh_sessions" USING btree("user_id");