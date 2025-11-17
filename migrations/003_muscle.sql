-- +goose Up

CREATE TYPE muscle_group_type AS ENUM (
    'chest', 
    'back', 
    'shoulders', 
    'arms',
    'legs',
    'core',
    'neck'
);

CREATE TABLE muscle (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    muscle_group muscle_group_type NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS muscle;
DROP TYPE IF EXISTS muscle_group_type;


