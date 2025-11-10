package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetInstructionByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetInstructionByIDResp struct {
	Instruction *exercise.Instruction
}

func (qry *GetInstructionByIDQuery) Handle(ctx context.Context) (any, error) {
	instruction, err := qry.Read.Exercise.Instruction.GetByID(ctx, qry.ID)
	if err != nil {
		return GetInstructionByIDResp{}, fmt.Errorf("failed to get exercise instruction: %w", err)
	}

	return GetInstructionByIDResp{Instruction: instruction}, nil
}
