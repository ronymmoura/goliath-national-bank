ALTER TABLE IF EXISTS "accounts" DROP COLUMN IF EXISTS "user_id";

DROP TABLE IF EXISTS "users";

ALTER TABLE IF EXISTS "accounts" ADD COLUMN IF NOT EXISTS "owner" varchar NOT NULL;

