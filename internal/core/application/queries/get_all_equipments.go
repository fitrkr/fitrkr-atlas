package queries

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
)

type GetAllEquipmentsQuery struct{}

type GetAllEquipmentsResp struct {
	Equipments []*equipment.Equipment
}

func (qry *GetAllEquipmentsQuery) Handle(ctx context.Context) (any, error) {
	equipments, err := read.Equipment.GetAll(ctx)
	if err != nil {
		return GetAllEquipmentsResp{}, fmt.Errorf("failed to get equipments: %w", err)
	}

	return GetAllEquipmentsResp{Equipments: equipments}, nil
}

func init() {
	register(&GetAllEquipmentsQuery{})
}
