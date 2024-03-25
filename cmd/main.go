package main

import (
	"context"
	"flood-control/internal/config"
	"flood-control/internal/db/postgres"
	"flood-control/internal/services/floodcontrol"
)

const N = 5
const K = 5

const exampleUserId = 2

func main() {
	cfg := config.GetConfig()
	postgresClient, err := postgres.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	var floodControlService FloodControl = floodcontrol.NewFloodControlService(K, N, postgresClient)

	_, _ = floodControlService.Check(ctx, exampleUserId)
}

type FloodControl interface {
	Check(ctx context.Context, userID int64) (bool, error)
}
