-- +goose Up

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE ,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

CREATE TABLE subcategories (
    id SERIAL PRIMARY KEY,
    category_id INT NOT NULL REFERENCES categories(id) ON DELETE RESTRICT,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(name, category_id)
);

-- +goose Down
DROP TABLE IF EXISTS subcategories;
DROP TABLE IF EXISTS categories;
