package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetEquipmentByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetEquipmentByIDResp struct {
	Equipment *equipment.Equipment
}

func (qry *GetEquipmentByIDQuery) Handle(ctx context.Context) (any, error) {
	eq, err := qry.Read.Equipment.GetByID(ctx, qry.ID)
	if err != nil {
		return GetEquipmentByIDResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	return GetEquipmentByIDResp{Equipment: eq}, nil
}
