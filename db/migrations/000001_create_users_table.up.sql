
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY not null,
  "in_time" varchar not null,
  "in_meet" boolean not null default(false)
);