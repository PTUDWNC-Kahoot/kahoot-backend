CREATE TABLE "register_orders" (
  "id" SERIAL PRIMARY KEY,
  "email" varchar(255),
  "verify_code" int,
  "expires_at" timestamp
);

