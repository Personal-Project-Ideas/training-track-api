CREATE TABLE IF NOT EXISTS "training_logs"."registered_trainings" (
  "id" bigserial PRIMARY KEY,
  "user_id" uuid,
  "sport_id" uuid,
  "intensity" integer,
  "duration" integer,
  "training_date" date,
  "created_at" timestamptz,
  "updated_at" timestamptz
);

ALTER TABLE "training_logs"."registered_trainings"
  ADD CONSTRAINT "fk_registered_trainings_user"
  FOREIGN KEY ("user_id") REFERENCES "training_logs"."users" ("id");

ALTER TABLE "training_logs"."registered_trainings"
  ADD CONSTRAINT "fk_registered_trainings_sport"
  FOREIGN KEY ("sport_id") REFERENCES "training_logs"."sports" ("id");
-- rollback
-- DROP TABLE IF EXISTS "training_logs"."registered_trainings";

