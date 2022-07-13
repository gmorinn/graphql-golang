BEGIN;

CREATE TABLE "files" (
    "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
    "created_at" timestamptz NOT NULL DEFAULT (NOW()),
    "updated_at" timestamptz NOT NULL DEFAULT (NOW()),
    "deleted_at" timestamptz,
    "name" text,
    "url" text,
    "mime" text,
    "size" bigint
);

COMMIT;