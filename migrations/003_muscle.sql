-- +goose Up

CREATE TABLE muscle_group (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

CREATE TABLE muscle (
    id SERIAL PRIMARY KEY,
    muscle_group_id INT NOT NULL REFERENCES muscle_group(id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(name, muscle_group_id)
);

-- +goose Down
DROP TABLE IF EXISTS muscle;
DROP TABLE IF EXISTS muscle_group;
