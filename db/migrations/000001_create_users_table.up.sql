CREATE TABLE "meetings_1" (
  "id" SERIAL not null,
  "meeting_room" int ,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false));

CREATE TABLE "meetings_2" (
  "id" SERIAL not null,
  "meeting_room" int ,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false));

insert into "meetings_1"(id,meeting_room,in_time,in_meet)
values(1,1,'11:00','false'),
      (2,1,'11:30','false'),
      (3,1,'12:00','false'),
      (4,1,'12:30','false'),
      (5,1,'13:00','false'),
      (6,1,'13:30','false'),
      (7,1,'14:00','false'),
      (8,1,'14:30','false'),
      (9,1,'15:00','false'),
      (10,1,'15:30','false'),
      (11,1,'16:00','false'),
      (12,1,'16:30','false'),
      (13,1,'17:00','false'),
      (14,1,'17:30','false'),
      (15,1,'18:00','false'),
      (16,1,'18:30','false'),
      (17,1,'19:00','false');

insert into "meetings_2"(id,meeting_room,in_time,in_meet)
values(1,2,'11:00','false'),
      (2,2,'11:30','false'),
      (3,2,'12:00','false'),
      (4,2,'12:30','false'),
      (5,2,'13:00','false'),
      (6,2,'13:30','false'),
      (7,2,'14:00','false'),
      (8,2,'14:30','false'),
      (9,2,'15:00','false'),
      (10,2,'15:30','false'),
      (11,2,'16:00','false'),
      (12,2,'16:30','false'),
      (13,2,'17:00','false'),
      (14,2,'17:30','false'),
      (15,2,'18:00','false'),
      (16,2,'18:30','false'),
      (17,2,'19:00','false');