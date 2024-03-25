package postgres

import (
	"context"
	"flood-control/internal/db/model"
)

func (repo *Repository) GetUserCountById(ctx context.Context, userID uint64) (uint64, error) {
	requestToDb := `
		SELECT user_id, calls_count FROM user_calls WHERE user_id = $1
	`

	var userCalls model.UserCalls
	row := repo.client.QueryRow(ctx, requestToDb, userID)
	err := row.Scan(&userCalls.UserId, &userCalls.CallsCount)
	if err != nil {
		return 0, err
	}

	return userCalls.CallsCount, nil
}
