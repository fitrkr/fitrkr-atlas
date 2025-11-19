-- +goose Up

CREATE TABLE exercise_view (
    id INT PRIMARY KEY REFERENCES exercise(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    difficulty VARCHAR(50) NOT NULL,
    position VARCHAR(50) NOT NULL,
    alias TEXT[],
    equipment JSONB,
    muscle JSONB,
    category JSONB,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, -- Soft delete 
    purge_at TIMESTAMP -- Auto-delete after 45 days of soft delete
);

-- +goose Down
DROP TABLE IF EXISTS exercise_view;
