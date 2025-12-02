package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type UpdateAlias struct {
	Add    []exercise.Alias `json:"add,omitempty"`
	Remove []int            `json:"remove,omitempty"`
}

type UpdateMuscle struct {
	Add    []MuscleReq `json:"add,omitempty"`
	Remove []int       `json:"remove,omitempty"`
}

type UpdateCategory struct {
	Add    []int `json:"add,omitempty"`
	Remove []int `json:"remove,omitempty"`
}

type UpdateAttachment struct {
	Add    []int `json:"add,omitempty"`
	Remove []int `json:"remove,omitempty"`
}

type UpdateExerciseCommand struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Difficulty  string            `json:"difficulty"`
	Position    string            `json:"position"`
	EquipmentID *int              `json:"equipment_id,omitempty"`
	Aliases     *UpdateAlias      `json:"aliases,omitempty"`
	Muscles     *UpdateMuscle     `json:"muscles,omitempty"`
	Categories  *UpdateCategory   `json:"categories,omitempty"`
	Attachments *UpdateAttachment `json:"attachments,omitempty"`
}

type UpdateExerciseResp struct {
	Exercise exercise.Exercise
}

func (cmd *UpdateExerciseCommand) Handle(ctx context.Context) (any, error) {
	exists, err := read.Exercise.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, ports.ErrExerciseNotFound
	}

	if cmd.Name != "" {
		cmd.Name = strings.ToLower(strings.TrimSpace(cmd.Name))

		// Check if an alias exists with the same name
		if _, err := read.Exercise.Alias.GetByName(ctx, cmd.Name); err == nil {
			return nil, fmt.Errorf("alias already exists")
		}

		exists.Name = cmd.Name
	}

	if cmd.Description != "" {
		description := strings.ToLower(cmd.Description)
		exists.Description = &description
	}

	if cmd.Difficulty != "" {
		difficulty, err := exercise.NewDifficulty(cmd.Difficulty)
		if err != nil {
			return UpdateExerciseResp{}, exercise.ErrInvalidDifficulty
		}
		exists.Difficulty = difficulty.ToString()
	}

	if cmd.Position != "" {
		position, err := exercise.NewBodyPosition(cmd.Position)
		if err != nil {
			return UpdateExerciseResp{}, exercise.ErrInvalidPosition
		}
		exists.Position = position.ToString()
	}

	// Validate equipment if provided
	if cmd.EquipmentID != exists.EquipmentID {
		// if it's not nil just verify it exists
		if cmd.EquipmentID != nil {
			_, err := read.Equipment.GetByID(ctx, *cmd.EquipmentID)
			if err != nil {
				return UpdateExerciseResp{}, fmt.Errorf("failed to read equipment %d: %w", *cmd.EquipmentID, err)
			}
		}
		// if the EquipmentID is nil still assign it
		exists.EquipmentID = cmd.EquipmentID
	}

	exists.Touch()

	ex, err := write.Exercise.Update(ctx, *exists)
	if err != nil {
		return nil, fmt.Errorf("failed to update exercise: %w", err)
	}

	if err = cmd.build(ctx); err != nil {
		return UpdateExerciseResp{}, fmt.Errorf("failed to update exercise relations: %w", err)
	}

	return UpdateExerciseResp{Exercise: ex}, nil
}

func (cmd *UpdateExerciseCommand) build(ctx context.Context) error {
	builder := NewExerciseBuilder(ctx, cmd.ID)

	if cmd.EquipmentID != nil {
		builder.equipmentID = cmd.EquipmentID
	}

	if cmd.Aliases != nil {
		builder.WithAlias(cmd.Aliases.Add, cmd.Aliases.Remove)
	}

	if cmd.Muscles != nil {
		builder.WithMuscles(cmd.Muscles.Add, cmd.Muscles.Remove)
	}

	if cmd.Categories != nil {
		builder.WithCategories(cmd.Categories.Add, cmd.Categories.Remove)
	}

	if cmd.Attachments != nil {
		builder.WithAttachments(cmd.Attachments.Add, cmd.Attachments.Remove)
	}

	if err := builder.Execute(); err != nil {
		return err
	}

	return nil
}

func init() {
	register(&UpdateExerciseCommand{})
}
