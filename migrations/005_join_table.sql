-- +goose Up

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
    activation_level VARCHAR(20) NOT NULL, 
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, muscle_id)
);

CREATE TABLE exercise_equipment (
    id INT PRIMARY KEY REFERENCES exercise(id) ON DELETE CASCADE,
    equipment_id INT NOT NULL REFERENCES equipment(id) ON DELETE RESTRICT,
    attachment_id INT REFERENCES equipment_attachment(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP 
);

CREATE TABLE exercise_category (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    subcategory_id INT NOT NULL REFERENCES subcategory(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, subcategory_id)
);

-- Different level of the same exercise
CREATE TABLE exercise_variation (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    variation_exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, variation_exercise_id)
);

-- Same level of the same exercise
CREATE TABLE exercise_alternative (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    alternate_exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, alternate_exercise_id)
);

CREATE TABLE exercise_media (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    media_type VARCHAR(20) NOT NULL, 
    display_order INT NOT NULL,
    is_primary BOOLEAN,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, url)
);

CREATE TABLE exercise_instruction (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercise(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    instruction_order INT NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, instruction_order)
);

-- +goose Down
DROP TABLE IF EXISTS exercise_instruction;
DROP TABLE IF EXISTS exercise_media;
DROP TABLE IF EXISTS exercise_alternative;
DROP TABLE IF EXISTS exercise_variation;
DROP TABLE IF EXISTS exercise_category;
DROP TABLE IF EXISTS exercise_equipment;
DROP TABLE IF EXISTS exercise_muscle;
DROP TABLE IF EXISTS exercise_alias;
