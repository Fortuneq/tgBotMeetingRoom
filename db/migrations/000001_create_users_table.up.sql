
CREATE TABLE "meetings" (
  "id" SERIAL not null,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false)
);

