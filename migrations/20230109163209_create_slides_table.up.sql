CREATE TABLE "slides" (
	"id" SERIAL PRIMARY KEY,
  "presentation_id" int,
  "type" int,
  "question" text,
  "time_limit" int,
  "points" int,
  "image_url" text,
  "video_url" text,
  "title" text,
  "text" text
);