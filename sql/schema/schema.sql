CREATE EXTENSION pgcrypto;

CREATE TYPE "role" AS ENUM (
  'admin',
  'pro',
  'user'
);

CREATE TABLE "students" (
  "id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT (now()),
  "deleted_at" timestamp CONSTRAINT deletedchk CHECK (deleted_at > created_at),
  "email" text NOT NULL CONSTRAINT emailchk CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
  "name" text CONSTRAINT namechk CHECK (char_length(name) >= 3 AND char_length(name) <= 20 AND  name ~ '^[^0-9]*$') DEFAULT NULL,
  "role" role NOT NULL DEFAULT 'user'
);
