-- +goose Up

CREATE TABLE muscle_groups (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

CREATE TABLE muscles (
    id SERIAL PRIMARY KEY,
    muscle_group_id INT NOT NULL REFERENCES muscle_groups(id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(name, muscle_group_id)
);

-- +goose Down
DROP TABLE IF EXISTS muscles;
DROP TABLE IF EXISTS muscle_groups;
