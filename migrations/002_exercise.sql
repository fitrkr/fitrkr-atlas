-- +goose Up

CREATE TYPE difficulty_level AS ENUM ('beginner', 'intermediate', 'advanced', 'elite');
CREATE TYPE body_position AS ENUM (
   'standing', 
   'sitting',   
   'kneeling',  
   'prone',     
   'supine',    
   'sidelying',
   'quadruped',
   'halfkneeling',
   'inverted',  
   'hanging'
);

-- Main Table 
CREATE TABLE exercise (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    difficulty difficulty_level NOT NULL,
    position body_position NOT NULL,
    equipment_id INT REFERENCES equipment(id),
    created_at TIMESTAMP, 
    updated_at TIMESTAMP, 
    deleted_at TIMESTAMP, -- Soft delete 
    purge_at TIMESTAMP -- Auto-delete after 45 days of soft delete
);

-- +goose Down
DROP TABLE IF EXISTS exercise;
DROP TYPE IF EXISTS difficulty_level;
DROP TYPE IF EXISTS exercise_type;
DROP TYPE IF EXISTS body_position;
