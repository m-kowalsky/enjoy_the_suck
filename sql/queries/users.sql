-- name: GetUserById :one
Select * from users
where id = (?);

-- name: GetUserByEmail :one
Select * from users
where email = (?)
Order by id;

-- name: CreateUser :one
Insert into users (id, created_at, updated_at, email)
Values ( ?, ?, ?, ?)
Returning *;
