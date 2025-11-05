-- +goose Up

CREATE TABLE equipments (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE ,
    description TEXT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP
);

CREATE TABLE equipment_attachments (
    id SERIAL PRIMARY KEY,
    equipment_id INT NOT NULL REFERENCES equipments(id) ON DELETE RESTRICT,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(name, equipment_id)
);

-- +goose Down
DROP TABLE IF EXISTS equipment_attachments;
DROP TABLE IF EXISTS equipments;
