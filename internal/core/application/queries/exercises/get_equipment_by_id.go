package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetEquipmentByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetEquipmentByIDResp struct {
	Equipment *exercise.ExerciseEquipment
}

func (qry *GetEquipmentByIDQuery) Handle(ctx context.Context) (any, error) {
	equipment, err := qry.Read.Exercise.Equipment.GetByID(ctx, qry.ID)
	if err != nil {
		return GetEquipmentByIDResp{}, fmt.Errorf("failed to get exercise equipment: %w", err)
	}

	return GetEquipmentByIDResp{Equipment: equipment}, nil
}
