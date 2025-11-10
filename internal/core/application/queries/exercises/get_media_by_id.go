package exercises

import (
	"context"
	"fmt"

	"github.com/cheezecakee/fitrkr-atlas/internal/core/domain/exercise"
	"github.com/cheezecakee/fitrkr-atlas/internal/core/ports"
)

type GetMediaByIDQuery struct {
	ID   int `json:"id"`
	Read ports.Read
}

type GetMediaByIDResp struct {
	Media *exercise.Media
}

func (qry *GetMediaByIDQuery) Handle(ctx context.Context) (any, error) {
	media, err := qry.Read.Exercise.Media.GetByID(ctx, qry.ID)
	if err != nil {
		return GetMediaByIDResp{}, fmt.Errorf("failed to get exercise media: %w", err)
	}

	return GetMediaByIDResp{Media: media}, nil
}
