package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type CreateAlias struct {
	Add []exercise.Alias `json:"add,omitempty"`
}

type CreateMuscle struct {
	Add []MuscleReq `json:"add,omitempty"`
}

type CreateExerciseCommand struct {
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Difficulty  string        `json:"difficulty"`
	Position    string        `json:"position"`
	Type        string        `json:"type"`
	EquipmentID *int          `json:"equipment_id,omitempty"`
	Aliases     *CreateAlias  `json:"aliases,omitempty"`
	Muscles     *CreateMuscle `json:"muscles,omitempty"`
	Categories  []int         `json:"categories,omitempty"`
	Attachments []int         `json:"attachments,omitempty"`
}

type CreateExerciseResp struct {
	ID int
}

func (cmd *CreateExerciseCommand) Handle(ctx context.Context) (any, error) {
	// Check if exercise already exists
	name := strings.TrimSpace(strings.ToLower(cmd.Name))
	if _, err := read.Exercise.GetByName(ctx, name); err == nil {
		return nil, ports.ErrDuplicateExercise
	}

	// Check if alias exists with the same name
	if _, err := read.Exercise.Alias.GetByName(ctx, name); err == nil {
		return nil, ports.ErrDuplicateExercise
	}

	// Validate difficulty
	difficulty, err := exercise.NewDifficulty(cmd.Difficulty)
	if err != nil {
		return CreateExerciseResp{}, exercise.ErrInvalidDifficulty
	}

	// Validate body position
	position, err := exercise.NewBodyPosition(cmd.Position)
	if err != nil {
		return CreateExerciseResp{}, exercise.ErrInvalidPosition
	}

	// Create Exercise
	ex, err := exercise.New(name, cmd.Description, difficulty.ToString(), position.ToString())
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to create new exercise: %w", err)
	}

	// Validate equipment if provided
	if cmd.EquipmentID != nil {
		equipment, err := read.Equipment.GetByID(ctx, *cmd.EquipmentID)
		if err != nil {
			return CreateExerciseResp{}, fmt.Errorf("failed to read equipment %d: %w", *cmd.EquipmentID, err)
		}
		ex.EquipmentID = equipment.ID
	}

	// Add exercise to DB and get exercise back
	ex, err = write.Exercise.Add(ctx, ex)
	if err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to insert exercise: %w", err)
	}

	if err = cmd.build(ctx, *ex.ID); err != nil {
		return CreateExerciseResp{}, fmt.Errorf("failed to create exercise relations: %w", err)
	}

	return CreateExerciseResp{ID: *ex.ID}, nil
}

func (cmd *CreateExerciseCommand) build(ctx context.Context, exerciseID int) error {
	builder := NewExerciseBuilder(ctx, exerciseID)

	if cmd.EquipmentID != nil {
		builder.equipmentID = cmd.EquipmentID
	}

	if cmd.Aliases != nil {
		builder.WithAlias(cmd.Aliases.Add, nil)
	}

	if cmd.Muscles != nil {
		builder.WithMuscles(cmd.Muscles.Add, nil)
	}

	if cmd.Categories != nil {
		builder.WithCategories(cmd.Categories, nil)
	}

	if cmd.Attachments != nil {
		builder.WithAttachments(cmd.Attachments, nil)
	}

	if err := builder.Execute(); err != nil {
		return err
	}

	return nil
}

func init() {
	register(&CreateExerciseCommand{})
}
