CREATE TABLE "meetings" (
  "id" SERIAL not null,
  "comment" text,
  "in_time" text PRIMARY KEY not null,
  "in_meet" boolean not null default(false));

insert into "meetings"(id,comment,in_time,in_meet)
values(1,'zoom','11:00','false'),
      (2,'backend','11:30','false'),
      (3,'proc','12:00','false'),
      (4,'webstorm','12:30','false'),
      (5,'project','13:00','false'),
      (6,'crypto','13:30','false'),
      (7,'frontend','14:00','false'),
      (8,'badbot','14:30','false'),
      (9,'fire of vlad','15:00','false'),
      (10,'hire of new dev','15:30','false'),
      (11,'','16:00','false'),
      (12,'','16:30','false'),
      (13,'','17:00','false'),
      (14,'','17:30','false'),
      (15,'','18:00','false'),
      (16,'','18:30','false'),
      (17,'','19:00','false');
