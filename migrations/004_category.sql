-- +goose Up

CREATE TYPE category_type AS ENUM (
    'strength', 
    'cardio', 
    'flexibility'
);

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE ,
    type category_type NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

-- +goose Down
DROP TABLE IF EXISTS category;
DROP TYPE IF EXISTS category_type;
