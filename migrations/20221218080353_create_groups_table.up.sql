CREATE TABLE "groups" (
  "id" SERIAL PRIMARY KEY,
  "admin_id" int,
  "name" varchar(50),
  "invitation_link" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);