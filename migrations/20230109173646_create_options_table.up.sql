CREATE TABLE "options" (
	"id" SERIAL PRIMARY KEY,
  "slide_id" int,
	"image_url" text,
	"color" text,
	"content" text,
	"is_correct" boolean
);
