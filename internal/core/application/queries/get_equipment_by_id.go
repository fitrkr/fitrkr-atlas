package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr/atlas/internal/core/domain/equipment"
)

type GetEquipmentByIDQuery struct {
	ID int `json:"id"`
}

type GetEquipmentByIDResp struct {
	Equipment *equipment.Equipment
}

func (qry *GetEquipmentByIDQuery) Handle(ctx context.Context) (any, error) {
	eq, err := read.Equipment.GetByID(ctx, qry.ID)
	if err != nil {
		return GetEquipmentByIDResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	return GetEquipmentByIDResp{Equipment: eq}, nil
}

func init() {
	register(&GetEquipmentByIDQuery{})
}
