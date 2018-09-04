--- +goose Up
CREATE TABLE users (
	id SERIAL,
	name varchar(50),
	username varchar(30) NOT NULL,
	password varchar(255) NOT NULL,
	created_at timestamptz,
	update_at timestamptz,
	deleted_at timestamptz,
	created_by timestamptz,
	updated_by timestamptz,
	deleted_by timestamptz,
	PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE user;
