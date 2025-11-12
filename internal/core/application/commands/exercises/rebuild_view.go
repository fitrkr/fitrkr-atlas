package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/view"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type ViewBuilder struct {
	Write ports.Write
	Read  ports.Read
}

func (vb *ViewBuilder) RebuildView(ctx context.Context, exerciseID int) error {
	ex, err := vb.Read.Exercise.GetByID(ctx, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to read exercise: %w", err)
	}

	equipmentView, err := vb.getEquipmentView(ctx, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to get equipment view: %w", err)
	}

	muscleGroupView, err := vb.getMuscleGroupView(ctx, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to get muscle view: %w", err)
	}

	categoryView, err := vb.getCategoryView(ctx, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to get category view: %w", err)
	}

	aliases, err := vb.getAliasView(ctx, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to get aliases: %w", err)
	}

	// Get existing view and replace it with updates without touching media and instruction
	existing, err := vb.Read.View.GetByID(ctx, exerciseID)
	if err != nil {
		return fmt.Errorf("failed to read view: %w", err)
	}

	existing.Name = ex.Name
	existing.Description = *ex.Description
	existing.Type = ex.Type.ToString()
	existing.Difficulty = ex.Difficulty.ToString()
	existing.BodyPosition = string(ex.BodyPosition)
	existing.MuscleGroup = muscleGroupView
	existing.Equipment = *equipmentView
	existing.Category = categoryView
	existing.Alias = aliases
	existing.CreatedAt = ex.CreatedAt
	existing.UpdatedAt = ex.UpdatedAt
	existing.DeletedAt = ex.DeletedAt
	existing.PurgeAt = ex.PurgedAt

	err = vb.Write.View.Update(ctx, *existing)
	if err != nil {
		return fmt.Errorf("failed to update view: %w", err)
	}

	return nil
}

func (vb *ViewBuilder) getEquipmentView(ctx context.Context, exerciseID int) (*view.Equipment, error) {
	exEquipment, err := vb.Read.Exercise.Equipment.GetByID(ctx, exerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to read equipment: %w", err)
	}

	if exEquipment == nil {
		return nil, fmt.Errorf("exercise equipment cannot empty")
	}

	attachment := view.Attachment{}

	equipment, err := vb.Read.Equipment.GetByID(ctx, exEquipment.EquipmentID)
	if err != nil {
		return nil, fmt.Errorf("failed to read equipment %d: %w", exEquipment.EquipmentID, err)
	}

	if exEquipment.AttachmentID != nil {
		// If there is an attachment, validate it
		a, err := vb.Read.Attachment.GetByID(ctx, *exEquipment.AttachmentID)
		if err != nil {
			return nil, fmt.Errorf("failed to read attachment %d: %w", exEquipment.AttachmentID, err)
		}
		attachment.ID = *exEquipment.AttachmentID
		attachment.Name = a.Name
		attachment.Type = a.Type.ToString()
	}

	return &view.Equipment{
		ID:          *equipment.ID,
		Name:        equipment.Name,
		Description: *equipment.Description,
		Type:        equipment.Type.ToString(),
		Attachment:  attachment,
	}, nil
}

func (vb *ViewBuilder) getMuscleGroupView(ctx context.Context, exerciseID int) ([]view.MuscleGroup, error) {
	muscleGroupView := []view.MuscleGroup{}

	exMuscles, err := vb.Read.Exercise.Muscle.GetByExerciseID(ctx, exerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise muscles: %w", err)
	}

	if len(exMuscles) == 0 {
		return muscleGroupView, nil
	}
	seen := make(map[int][]view.Muscle)

	for _, i := range exMuscles {
		m, err := vb.Read.Muscle.GetByID(ctx, i.MuscleID)
		if err != nil {
			return nil, ports.ErrMuscleNotFound
		}

		seen[m.MuscleGroupID] = append(seen[m.MuscleGroupID], view.Muscle{
			ID:         *m.ID,
			Name:       m.Name,
			Activation: i.ActivationLevel.ToString(),
		})
	}

	for muscleGroupID, muscle := range seen {
		mg, err := vb.Read.Muscle.Group.GetByID(ctx, muscleGroupID)
		if err != nil {
			return nil, fmt.Errorf("failed to read muscle group: %w", err)
		}

		// Add to view grouping
		muscleGroupView = append(muscleGroupView, view.MuscleGroup{
			ID:          muscleGroupID,
			Name:        string(mg.Name),
			Description: *mg.Description,
			Muscle:      muscle,
		})
	}
	return muscleGroupView, nil
}

func (vb *ViewBuilder) getCategoryView(ctx context.Context, exerciseID int) ([]view.Category, error) {
	categoryView := []view.Category{}

	exCategory, err := vb.Read.Exercise.Category.GetByExerciseID(ctx, exerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise category: %w", err)
	}
	if len(exCategory) == 0 {
		return categoryView, nil
	}

	seen := make(map[int][]view.Subcategory)

	for _, i := range exCategory {
		subcategory, err := vb.Read.Category.Subcategory.GetByID(ctx, i.SubcategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to read subcategory: %w", err)
		}
		// Add to view grouping
		seen[subcategory.CategoryID] = append(seen[subcategory.CategoryID], view.Subcategory{
			ID:   *subcategory.ID,
			Name: subcategory.Name,
		})
	}

	for categoryID, subcategories := range seen {
		c, err := vb.Read.Category.GetByID(ctx, categoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to read category: %w", err)
		}

		categoryView = append(categoryView, view.Category{
			ID:          categoryID,
			Name:        string(c.Name),
			Subcategory: subcategories,
		})
	}
	return categoryView, nil
}

func (vb *ViewBuilder) getAliasView(ctx context.Context, exerciseID int) ([]view.Alias, error) {
	aliasView := []view.Alias{}

	aliases, err := vb.Read.Exercise.Alias.GetByExerciseID(ctx, exerciseID)
	if err != nil {
		return nil, fmt.Errorf("failed to read exercise aliases: %w", err)
	}

	if len(aliases) == 0 {
		return aliasView, nil
	}

	for _, i := range aliases {
		alias := view.Alias{
			ID:           *i.ID,
			Name:         i.Name,
			LanguageCode: i.LanguageCode,
		}

		aliasView = append(aliasView, alias)
	}

	return aliasView, nil
}
