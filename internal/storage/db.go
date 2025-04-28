package storage

import (
	"auth-service-medods/pkg/logger"
	"context"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	Psql *pgx.Conn
}

// Открываю коннект
func Open(url string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, logger.LogError("Database connection error: %w", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		conn.Close(context.Background())
		return nil, logger.LogError("Error checking database connection: %w", err)
	}

	logger.LogInfo("Successfully connected to database")
	return &DB{Psql: conn}, nil
}

// Закрываю соединение
func (db *DB) Close() {
	if db.Psql != nil {
		db.Psql.Close(context.Background())
		logger.LogInfo("Closed database connection")
	}
}
