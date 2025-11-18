-- name: GetUpcomingEvents :many
SELECT * FROM event WHERE time > CURRENT_TIMESTAMP;


-- name: GetEvents :many
SELECT * FROM event;

-- name: GetCathegorys :many
SELECT * FROM cathegory;

-- name: NewCathegory :one
INSERT INTO cathegory (name)
VALUES (?) RETURNING id;

-- name: NewEvent :one
INSERT INTO event (name, description, cathegory, time)
VALUES (?, ?, (SELECT id FROM cathegory WHERE cathegory.name = ?), ?) RETURNING id;

