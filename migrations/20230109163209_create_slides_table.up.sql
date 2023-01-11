CREATE TABLE "slides" (
	"id" SERIAL PRIMARY KEY,
  "presentation_id" int,
  "time_limit" int,
  "type" int,
  "question" text,
  "points" int,
  "heading" text,
	"sub_heading" text,    
	"paragraph" text,  
  "image_url" text
);