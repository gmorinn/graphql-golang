BEGIN;

ALTER TABLE "students"
    ADD COLUMN "password" text NOT NULL CONSTRAINT passwordchk CHECK (char_length(password) >= 9);

COMMIT;