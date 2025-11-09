-- +goose Up

CREATE TYPE difficulty_level AS ENUM ('beginner', 'intermediate', 'advanced', 'elite');
CREATE TYPE exercise_type AS ENUM ('body_weight', 'free_weight', 'machine');

-- Main Table 
CREATE TABLE exercise (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    type exercise_type NOT NULL,
    difficulty difficulty_level NOT NULL,
    body_position VARCHAR(50) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, -- Soft delete 
    purge_at TIMESTAMP -- Auto-delete after 45 days of soft delete
);

-- +goose Down
DROP TABLE IF EXISTS exercise;
DROP TYPE IF EXISTS difficulty_level;
DROP TYPE IF EXISTS exercise_type;
