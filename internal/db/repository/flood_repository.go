package repository

import (
	"context"
	"flood-control/internal/db/model"
)

type FloodRepository interface {
	SaveFlood(context.Context, model.Flood) error
	GetFloodByUserId(context.Context, uint64) (model.Flood, error)
	DeleteFloodByUserId(context.Context, uint64) error
}
