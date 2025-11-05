-- +goose Up

-- Search/Lookup
CREATE INDEX idx_exercises_name ON exercises(name);
CREATE INDEX idx_exercises_difficulty ON exercises(difficulty);

-- Join table lookups (for filtering)
CREATE INDEX idx_exercise_muscles_exercise_id ON exercise_muscles(exercise_id);
CREATE INDEX idx_exercise_muscles_muscle_id ON exercise_muscles(muscle_id);

CREATE INDEX idx_exercise_equipments_exercise_id ON exercise_equipments(exercise_id);
CREATE INDEX idx_exercise_equipments_equipment_id ON exercise_equipments(equipment_id);

CREATE INDEX idx_exercise_categories_exercise_id ON exercise_categories(exercise_id);
CREATE INDEX idx_exercise_categories_subcategory_id ON exercise_categories(subcategory_id);

-- Relationships (variations, alternatives)
CREATE INDEX idx_exercise_variations_base ON exercise_variations(base_exercise_id);
CREATE INDEX idx_exercise_variations_variation ON exercise_variations(variation_exercise_id);

CREATE INDEX idx_exercise_alternatives_base ON exercise_alternatives(base_exercise_id);
CREATE INDEX idx_exercise_alternatives_alt ON exercise_alternatives(alternate_exercise_id);

-- Soft deletes
CREATE INDEX idx_exercises_deleted ON exercises(deleted_at);

-- +goose Down

DROP INDEX IF EXISTS idx_exercises_name;
DROP INDEX IF EXISTS idx_exercises_difficulty;
DROP INDEX IF EXISTS idx_exercise_muscles_exercise_id;
DROP INDEX IF EXISTS idx_exercise_muscles_muscle_id;
DROP INDEX IF EXISTS idx_exercise_equipments_exercise_id;
DROP INDEX IF EXISTS idx_exercise_equipments_equipment_id;
DROP INDEX IF EXISTS idx_exercise_categories_exercise_id;
DROP INDEX IF EXISTS idx_exercise_categories_subcategory_id;
DROP INDEX IF EXISTS idx_exercise_variations_base;
DROP INDEX IF EXISTS idx_exercise_variations_variation;
DROP INDEX IF EXISTS idx_exercise_alternatives_base;
DROP INDEX IF EXISTS idx_exercise_alternatives_alt;
DROP INDEX IF EXISTS idx_exercises_deleted;
