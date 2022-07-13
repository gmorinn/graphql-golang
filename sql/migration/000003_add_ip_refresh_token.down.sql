BEGIN;

ALTER TABLE "refresh_token" DROP COLUMN IF EXISTS "ip";
ALTER TABLE "refresh_token" DROP COLUMN IF EXISTS  "user_agent";

COMMIT;