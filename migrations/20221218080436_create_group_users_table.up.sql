CREATE TABLE "group_users" (
  "group_id" int,
  "user_id" int,
  "role" int,
  "name" text,
  PRIMARY KEY ("group_id", "user_id")
);