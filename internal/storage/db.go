package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Psql *pgxpool.Pool
}

// открываю соединение
func Open(url string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, fmt.Errorf("error checking database connection: %w", err)
	}

	log.Println("successfully connected to database")
	return &DB{Psql: pool}, nil
}

// закрываю соединение
func (db *DB) Close() {
	if db.Psql != nil {
		db.Psql.Close()
		log.Println("Closed")
	}
}
