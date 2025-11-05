-- +goose Up

-- Main Table 
CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    difficulty VARCHAR(20) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, -- Soft delete 
    purge_at TIMESTAMP -- Auto-delete after 45 days of soft delete
);

-- +goose Down
DROP TABLE IF EXISTS exercises;
