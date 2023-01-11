CREATE TABLE "presentations" (
	"id" SERIAL PRIMARY KEY,
  "group_id" int,
  "owner" int,
  "title" varchar(50),
  "description" text,
  "cover_image_url" text,
  "visibility" boolean
);