CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar(255),
  "password" varchar(255),
  "name" varchar(255),
  "workplace" varchar(255),
  "organization" varchar(255),
  "cover_image_url" text,
  "players" int,
  "plays" int,
  "kahoots" int,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);