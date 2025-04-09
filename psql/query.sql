-- name: GetUser :one
SELECT * FROM users
WHERE ID = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY name;

-- name: CreateUser :one
INSERT INTO users (
  name, email, bio
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
  set name = $2,
  email = $3,
  bio = $4
WHERE ID = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE ID = $1;

-- name: GetAssignedTasks :many
SELECT * FROM tasks 
WHERE assigned_to_user_id = $1 
ORDER BY created_at DESC;

-- name: GetCreatedTasks :many
SELECT * FROM tasks
WHERE created_by_user_id = $1
ORDER BY created_at DESC;

-- name: GetTask :one
SELECT * FROM tasks
WHERE ID = $1 LIMIT 1;

-- name: CreateTask :one
INSERT INTO tasks (
    title, description, created_by_user_id, assigned_to_user_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: AssignTask :exec
UPDATE tasks
SET assigned_to_user_id = $2
WHERE ID = $1
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE ID = $1;