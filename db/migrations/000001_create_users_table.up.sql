
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY not null,
  "full_name" varchar not null,
  "created_at" timestamp not null default(now()),
  "in_meet" boolean not null default(false)
);