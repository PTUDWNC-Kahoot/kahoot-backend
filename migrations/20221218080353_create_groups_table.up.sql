CREATE TABLE "groups" (
  "id" SERIAL PRIMARY KEY,
  "owner" int,
  "name" varchar(50),
  "cover_image_url" text,
  "invitation_link" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);
