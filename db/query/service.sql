-- name: CreateService :one
INSERT INTO service (
  name,
  description,
  versions
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListServices :many
SELECT * FROM service
ORDER BY name
LIMIT $1
OFFSET $2;

-- name: GetService :one
SELECT * FROM service
WHERE name = $1 LIMIT 1;


-- name: UpdateService :one
UPDATE service
  set name = $2,
  description = $3
WHERE name = $1
RETURNING *;

-- name: DeleteService :exec
DELETE FROM service
WHERE name = $1;