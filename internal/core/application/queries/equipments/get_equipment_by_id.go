package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/logr"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetEquipmentByIDQuery struct {
	ID   int `json:"id"`
	Read ports.EquipmentRead
}

type GetEquipmentByIDResp struct {
	Equipment *equipment.Equipment
}

func (qry *GetEquipmentByIDQuery) Handle(ctx context.Context) (any, error) {
	eq, err := qry.Read.GetByID(ctx, qry.ID)
	if err != nil {
		logr.Get().Errorf("failed to get equipment: %v", err)
		return GetEquipmentByIDResp{}, fmt.Errorf("failed to get equipment: %w", err)
	}

	return GetEquipmentByIDResp{Equipment: eq}, nil
}
