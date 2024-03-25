package flood

import dbrep "flood-control/internal/db/repository"
import domainrep "flood-control/internal/domain/repository"

import (
	"context"
	"flood-control/internal/db/model"
	"flood-control/internal/domain/entity"
)

type Repository struct {
	domainrep.FloodRepository
}

type FRepository struct {
	dbrep.FloodRepository
}

type UCRepository struct {
	dbrep.UserCallsRepository
}

func (repo *Repository) SaveFlood(ctx context.Context, flood *entity.Flood) /*map[string]*/ error {
	// not implemented
	floodDbRepo := FRepository{}
	return floodDbRepo.SaveFlood(ctx, DomainToDto(flood))
}

func (repo *Repository) DeleteFlood(ctx context.Context, userID uint64) /*map[string]*/ error {
	floodDbRepo := FRepository{}
	return floodDbRepo.DeleteFloodByUserId(ctx, userID)
}

func (repo *Repository) GetFlood(ctx context.Context, userID uint64) (entity.Flood, error) {
	floodDbRepo := FRepository{}
	floodDto, err := floodDbRepo.GetFloodByUserId(ctx, userID)
	if err != nil {
		return entity.Flood{}, err
	}

	userCallsDbRepo := UCRepository{}
	callsCount, err := userCallsDbRepo.GetUserCallsCountById(ctx, userID)
	if err != nil {
		return entity.Flood{}, err
	}

	return DtoToDomain(floodDto, callsCount), nil
}

func DomainToDto(flood *entity.Flood) model.Flood {
	return model.Flood{
		UserID: flood.UserID,
		EndAt:  flood.EndAt,
	}
}

func DtoToDomain(flood model.Flood, callsCount uint64) entity.Flood {
	return entity.Flood{
		UserID:     flood.UserID,
		EndAt:      flood.EndAt,
		CallsCount: callsCount,
	}
}
