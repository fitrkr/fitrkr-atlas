-- +goose Up
CREATE TYPE activation_level_type AS ENUM ('primary', 'secondary', 'tertiary');

-- Exercises
CREATE TABLE exercise_alias (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    language_code VARCHAR(5),
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, name, language_code)
);

CREATE TABLE exercise_muscle (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    muscle_id INT NOT NULL REFERENCES muscle(id) ON DELETE RESTRICT,
    activation activation_level_type NOT NULL, 
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, muscle_id)
);

CREATE TABLE exercise_equipment (
    id INT PRIMARY KEY REFERENCES exercise(id) ON DELETE CASCADE,
    equipment_id INT NOT NULL REFERENCES equipment(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

CREATE TABLE exercise_attachment (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    attachment_id INT REFERENCES attachment(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

CREATE TABLE exercise_category (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    category_id INT NOT NULL REFERENCES category(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, category_id)
);

-- +goose Down
DROP TABLE IF EXISTS exercise_category;
DROP TABLE IF EXISTS exercise_equipment;
DROP TABLE IF EXISTS exercise_attachment;
DROP TABLE IF EXISTS exercise_muscle;
DROP TABLE IF EXISTS exercise_alias;
DROP TYPE IF EXISTS activation_level_type;
