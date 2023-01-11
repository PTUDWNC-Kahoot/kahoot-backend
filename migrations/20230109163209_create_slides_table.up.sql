CREATE TABLE "slides" (
	"id" SERIAL PRIMARY KEY,
  "presentation_id" int,
  "type" int,
  "question" text,
  "heading" text,
	"sub_heading" text,    
	"paragraph" text,  
  "image_url" text
);