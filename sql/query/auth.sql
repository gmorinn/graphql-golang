-- name: LoginUser :one
SELECT id, name, email, role FROM students
WHERE email = $1
AND password = crypt($2, password)
AND deleted_at IS NULL;

-- name: Signup :one
INSERT INTO students (email, password, name) 
VALUES ($1, crypt($2, gen_salt('bf')), $3)
RETURNING *;

-- name: CheckEmailExist :one
SELECT EXISTS(
    SELECT id, created_at, updated_at, deleted_at, email, password, name, role FROM students
    WHERE email = $1
    AND deleted_at IS NULL
);