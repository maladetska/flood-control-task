package repository

import (
	"context"
)

type UserCallsRepository interface {
	GetUserCallsCountById(context.Context, uint64) (uint64, error)
}
