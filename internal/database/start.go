package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTables(db *pgxpool.Pool) error {
	ctx := context.Background()

	_, err := db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id serial PRIMARY KEY,
			username VARCHAR(20) NOT NULL UNIQUE,
			email VARCHAR(70) NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS refresh_tokens (
			id serial PRIMARY KEY,
			user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			token TEXT NOT NULL,
			token_expires_at BIGINT NOT NULL
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
