package floodcontrol

import (
	"context"
	"flood-control/internal/db/postgres"
	"flood-control/internal/domain/repository"
	"sync"
	"time"
)

type FloodControlService struct {
	maxCallsCount uint64
	interval      time.Duration
	repo          repository.FloodRepository
	postgres      *postgres.Repository
	mu            sync.Mutex
}

func NewFloodControlService(maxCallsCount uint64, intervalBySeconds uint64, postgres *postgres.Repository) *FloodControlService {
	return &FloodControlService{
		maxCallsCount: maxCallsCount,
		interval:      time.Duration(intervalBySeconds) * time.Second,
		postgres:      postgres,
		mu:            sync.Mutex{},
	}
}

func (service *FloodControlService) Check(ctx context.Context, userID int64) (bool, error) {
	flood, err := service.repo.GetFlood(ctx, uint64(userID))
	if err != nil {
		return false, err
	}

	err = flood.CallMethod(service.interval, service.maxCallsCount)
	if err = service.repo.SaveFlood(ctx, &flood); err != nil {
		return false, err
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
