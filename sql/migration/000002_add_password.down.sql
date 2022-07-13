BEGIN;

ALTER TABLE "students"
    DROP COLUMN "password";

ALTER TABLE "students" DROP CONSTRAINT "passwordchk";

COMMIT;