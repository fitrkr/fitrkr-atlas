-- +goose Up

CREATE TYPE equipment_type AS ENUM ('body_weight', 'free_weight', 'machine');
CREATE TYPE attachment_type AS ENUM ('band', 'cable', 'plate');

CREATE TABLE equipment (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE ,
    description TEXT,
    type equipment_type NOT NULL, 
    created_at TIMESTAMP, 
    updated_at TIMESTAMP
);

CREATE TABLE attachment (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    type attachment_type NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

CREATE TABLE equipment_attachment (
    id SERIAL PRIMARY KEY,
    equipment_id INT NOT NULL REFERENCES equipment(id) ON DELETE RESTRICT,
    attachment_id INT NOT NULL REFERENCES attachment(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(equipment_id, attachment_id)
);

-- +goose Down
DROP TABLE IF EXISTS equipment_attachment;
DROP TABLE IF EXISTS attachment;
DROP TABLE IF EXISTS equipment;
DROP TYPE IF EXISTS equipment_type;
DROP TYPE IF EXISTS attachment_type;
