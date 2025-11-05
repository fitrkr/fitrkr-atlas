-- +goose Up

-- Exercises
CREATE TABLE exercise_aliases (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    alias_name VARCHAR(255) NOT NULL,
    language_code VARCHAR(5),
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, alias_name, language_code)
);

CREATE TABLE exercise_muscles (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    muscle_id INT NOT NULL REFERENCES muscles(id) ON DELETE RESTRICT,
    activation_level VARCHAR(20) NOT NULL, 
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, muscle_id)
);

CREATE TABLE exercise_equipments (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    equipment_id INT NOT NULL REFERENCES equipments(id) ON DELETE RESTRICT,
    attachment_id INT REFERENCES equipment_attachments(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, equipment_id, attachment_id)
);

CREATE TABLE exercise_categories (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    subcategory_id INT NOT NULL REFERENCES subcategories(id) ON DELETE RESTRICT,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, subcategory_id)
);

-- Different level of the same exercise
CREATE TABLE exercise_variations (
    id SERIAL PRIMARY KEY,
    base_exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    variation_exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(base_exercise_id, variation_exercise_id)
);

-- Same level of the same exercise
CREATE TABLE exercise_alternatives (
    id SERIAL PRIMARY KEY,
    base_exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    alternate_exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(base_exercise_id, alternate_exercise_id)
);

CREATE TABLE exercise_media (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    media_type VARCHAR(20) NOT NULL, 
    display_order INT NOT NULL,
    is_primary BOOLEAN,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, url)
);

CREATE TABLE exercise_instructions (
    id SERIAL PRIMARY KEY,
    exercise_id INT NOT NULL REFERENCES exercises(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    instruction_order INT NOT NULL,
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    UNIQUE(exercise_id, instruction_order)
);

-- +goose Down
DROP TABLE IF EXISTS exercise_instructions;
DROP TABLE IF EXISTS exercise_media;
DROP TABLE IF EXISTS exercise_alternatives;
DROP TABLE IF EXISTS exercise_variations;
DROP TABLE IF EXISTS exercise_categories;
DROP TABLE IF EXISTS exercise_equipments;
DROP TABLE IF EXISTS exercise_muscles;
DROP TABLE IF EXISTS exercise_aliases;
