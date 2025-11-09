package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"fmt"
	"log"
	"os"
	"time"
)

// InitDB Connect to database
func InitDB() (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("ошибка подключения к базе: %v", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Fatalf("база недоступна: %v", err)
		return nil, err
	}

	log.Println("✅ Успешное подключение к PostgreSQL")

	return pool, nil
}
