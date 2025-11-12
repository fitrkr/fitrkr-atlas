-- +goose Up
CREATE TABLE exercise_view (
    id INT PRIMARY KEY REFERENCES exercise(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    type exercise_type NOT NULL, -- already exists
    difficulty difficulty_level NOT NULL, -- already exists 
    body_position VARCHAR(50) NOT NULL,
    alias JSONB,
    muscle JSONB,
    equipment JSONB,
    category JSONB,
    -- variation JSON,
    -- alternative JSON,
    media JSONB,
    instruction JSONB,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, -- Soft delete 
    purge_at TIMESTAMP, -- Auto-delete after 45 days of soft delete
);

-- +goose Down
DROP exercise_view;
