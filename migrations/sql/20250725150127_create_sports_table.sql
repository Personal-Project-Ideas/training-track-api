CREATE TABLE IF NOT EXISTS "training_logs"."sports" (
  "id" uuid PRIMARY KEY,
  "modality" varchar,
  "created_at" timestamptz,
  "updated_at" timestamptz
);
-- rollback
-- DROP TABLE IF EXISTS "training_logs"."sports";