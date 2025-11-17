-- +goose Up

-- Search/Lookup
CREATE INDEX idx_exercise_name ON exercise(name);
CREATE INDEX idx_exercise_difficulty ON exercise(difficulty);

-- Join table lookups (for filtering)
CREATE INDEX idx_exercise_muscle_exercise_id ON exercise_muscle(exercise_id);
CREATE INDEX idx_exercise_muscle_muscle_id ON exercise_muscle(muscle_id);

CREATE INDEX idx_exercise_equipment_id ON exercise_equipment(id);
CREATE INDEX idx_exercise_equipment_equipment_id ON exercise_equipment(equipment_id);

CREATE INDEX idx_exercise_attachment_id ON exercise_attachment(id);
CREATE INDEX idx_exercise_attachment_attachment_id ON exercise_attachment(attachment_id);

CREATE INDEX idx_exercise_category_exercise_id ON exercise_category(exercise_id);
CREATE INDEX idx_exercise_category_category_id ON exercise_category(category_id);

-- Soft deletes
CREATE INDEX idx_exercise_deleted ON exercise(deleted_at);

-- +goose Down

DROP INDEX IF EXISTS idx_exercise_name;
DROP INDEX IF EXISTS idx_exercise_difficulty;
DROP INDEX IF EXISTS idx_exercise_muscle_exercise_id;
DROP INDEX IF EXISTS idx_exercise_muscle_muscle_id;
DROP INDEX IF EXISTS idx_exercise_equipment_id;
DROP INDEX IF EXISTS idx_exercise_equipment_equipment_id;
DROP INDEX IF EXISTS idx_exercise_attachment_id;
DROP INDEX IF EXISTS idx_exercise_attachment_attachment_id;
DROP INDEX IF EXISTS idx_exercise_category_exercise_id;
DROP INDEX IF EXISTS idx_exercise_category_category_id;
DROP INDEX IF EXISTS idx_exercise_deleted;
