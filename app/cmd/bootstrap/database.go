package bootstrap

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/mkafonso/goledger-challenge-besu/config"
	dbrepository "github.com/mkafonso/goledger-challenge-besu/infra/db"
	db "github.com/mkafonso/goledger-challenge-besu/infra/db/sqlc"
)

func MustInitDatabase(ctx context.Context) *dbrepository.StorageRepository {
	conn, err := pgx.Connect(ctx, config.Env.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Fatal("failed to ping database:", err)
	}

	queries := db.New(conn)
	return dbrepository.NewStorageRepository(queries)
}
