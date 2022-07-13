BEGIN;

ALTER TABLE "refresh_token" ADD COLUMN "ip" text NOT NULL;
ALTER TABLE "refresh_token" ADD COLUMN "user_agent" text NOT NULL;

COMMIT;