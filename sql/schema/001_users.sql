-- +goose Up
Create table users (
id uuid primary key,
created_at timestamp not null,
updated_at timestamp not null,
email text not null unique
);

-- +goose Down
Drop table users;

