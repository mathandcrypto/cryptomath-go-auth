DROP TABLE IF EXISTS refresh_sessions;
ALTER TABLE refresh_sessions DROP CONSTRAINT IF EXISTS refresh_sessions_pkey;
DROP INDEX IF EXISTS user_idx;