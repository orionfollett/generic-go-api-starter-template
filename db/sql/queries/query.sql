-- name: GetRun :one
SELECT * FROM runs
WHERE id = $1 LIMIT 1;

-- name: CreateRun :one
INSERT INTO runs (
  title
) VALUES (
  $1
)
RETURNING *;