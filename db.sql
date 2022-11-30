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

CREATE TABLE "kahoots" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "title" varchar(255),
  "description" text,
  "cover_image_url" text,
  "visibility" boolean
);

CREATE TABLE "slides" (
  "id" SERIAL PRIMARY KEY,
  "kahoot_id" int,
  "type" int,
  "order" int,
  "question" text,
  "time_limit" int,
  "points" int,
  "image_url" text,
  "video_url" text,
  "answer_options" text,
  "title" varchar(255),
  "text" text
);

CREATE TABLE "answers" (
  "id" SERIAL PRIMARY KEY,
  "kahoot_id" int,
  "image_url" text,
  "color" int,
  "content" text,
  "is_correct" boolean,
  "order" int
);

CREATE TABLE "reports" (
  "game_id" int
);

CREATE TABLE "points" (
  "user_id" int,
  "kahoot_id" int,
  "turn_code" int,
  "nickname" varchar(50),
  "points" int
);

CREATE TABLE "groups" (
  "id" SERIAL PRIMARY KEY,
  "admin_id" int,
  "name" varchar(50),
  "invitation_link" text,
  "created_at" timestamp,
  "updated_at" timestamp,
  "deleted_at" timestamp
);

CREATE TABLE "group_kahoots" (
  "couse_id" int,
  "kahoot_id" int,
  "status" boolean
);

CREATE TABLE "group_users" (
  "group_id" int,
  "user_id" int,
  "role" int
  "name" text
);
