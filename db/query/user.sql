-- name: CreateMeeting :one
INSERT INTO meetings (
  in_time,
  in_meet
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetMeeting :one
SELECT * FROM meetings
WHERE id = $1 LIMIT 1;

-- name: GetMeetingForUpdate :one
SELECT * FROM meetings
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListMeetings :many
SELECT * FROM meetings
ORDER BY id
LIMIT $1
OFFSET $2;
