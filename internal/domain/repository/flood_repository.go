package repository

import (
	"context"
	"flood-control/internal/domain/entity"
)

type FloodRepository interface {
	SaveFlood(context.Context, *entity.Flood) error
	DeleteFlood(context.Context, uint64) error
	GetFlood(context.Context, uint64) (entity.Flood, error)
}
