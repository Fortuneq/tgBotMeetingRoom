
CREATE TABLE "meetings" (
  "id" SERIAL PRIMARY KEY not null,
  "in_time" text not null,
  "in_meet" boolean not null default(false)
);

insert into "meetings"(id,in_time,in_meet)
values(1,'11:00','false'),
      (2,'11:30','false'),
      (3,'12:00','false'),
      (4,'12:30','false'),
      (5,'13:00','false'),
      (6,'13:30','false'),
      (7,'14:00','false'),
      (8,'14:30','false'),
      (9,'15:00','false'),
      (10,'15:30','false'),
      (11,'16:00','false'),
      (12,'16:30','false'),
      (13,'17:00','false'),
      (14,'17:30','false'),
      (15,'18:00','false'),
      (16,'18:30','false'),
      (17,'19:00','false');