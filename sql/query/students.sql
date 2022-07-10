-- name: UpdateDescriptionStudent :exec
UPDATE students
SET name = $1,
    email = $2,
    updated_at = NOW()
WHERE id = $3;

-- name: GetStudentByID :one
SELECT * FROM students
WHERE id = $1
AND deleted_at IS NULL
LIMIT 1;

-- name: DeleteStudentByID :exec
UPDATE
    students
SET
    deleted_at = NOW()
WHERE 
    id = $1;

-- name: CountStudent :one
SELECT COUNT(*) FROM students
WHERE deleted_at IS NULL;

-- name: Liststudents :many
SELECT * FROM students
WHERE deleted_at IS NULL
AND (name ILIKE $1 OR email ILIKE $1)
LIMIT 5;
