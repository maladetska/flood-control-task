package postgres

import (
	"context"
	"errors"
	"flood-control/internal/db/model"
	"github.com/jackc/pgx/v5/pgconn"
)

func (repo *Repository) SaveFlood(ctx context.Context, floodPtr *model.Flood) error {
	request := `
		INSERT INTO flood (user_id, end_at) 
		VALUES ($1, $2)
		RETURNING user_id
	`

	row := repo.client.QueryRow(ctx, request, floodPtr.UserID, floodPtr.EndAt)
	if err := row.Scan(&floodPtr.UserID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			errors.As(err, &pgErr)
			return pgErr
		}
		return err
	}

	return nil
}

func (repo *Repository) GetFloodByUserId(ctx context.Context, userID uint64) (model.Flood, error) {
	request := `
		SELECT user_id, calls_count FROM user_calls WHERE user_id = $1
	`

	var flood model.Flood
	row := repo.client.QueryRow(ctx, request, userID)
	err := row.Scan(&flood.UserID, &flood.EndAt)
	if err != nil {
		return model.Flood{}, err
	}

	return flood, nil
}

func (repo *Repository) DeleteFloodByUserId(ctx context.Context, userID uint64) error {
	request := `
		DELETE FROM user_calls 
		       WHERE user_id = $1
		       RETURNING user_id
	`

	var flood model.Flood
	row := repo.client.QueryRow(ctx, request, userID)
	if err := row.Scan(&flood.UserID); err != nil {
		return err
	}

	return nil
}
