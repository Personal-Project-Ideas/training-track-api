CREATE TABLE IF NOT EXISTS "training_logs"."users" (
  "id" uuid PRIMARY KEY,
  "full_name" varchar(255),
  "age" integer,
  "email" varchar,
  "password" varchar,
  "created_at" timestamptz,
  "update_at" timestamptz
);

-- rollback
-- DROP TABLE IF EXISTS "training_logs"."users";

