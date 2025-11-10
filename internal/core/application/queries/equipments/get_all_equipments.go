// Package equipments
package equipments

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/equipment"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetAllEquipmentsQuery struct {
	Read ports.Read
}

type GetAllEquipmentsResp struct {
	Equipments []equipment.Equipment
}

func (qry *GetAllEquipmentsQuery) Handle(ctx context.Context) (any, error) {
	equipments, err := qry.Read.Equipment.GetAll(ctx)
	if err != nil {
		return GetAllEquipmentsResp{}, fmt.Errorf("failed to get equipments: %w", err)
	}

	return GetAllEquipmentsResp{Equipments: equipments}, nil
}
