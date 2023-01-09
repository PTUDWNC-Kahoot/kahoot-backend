CREATE TABLE "presentations" (
	"id" SERIAL PRIMARY KEY,
  "user_id" int,
  "title" varchar(50),
  "description" text,
  "cover_image_url" text,
  "visibility" boolean
);