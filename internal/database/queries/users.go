package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(db *pgxpool.Pool, username, email, password string) error {
	_, err := db.Exec(context.Background(), `
			INSERT INTO users (username, email, password)
			VALUES ($1, $2, $3)
		`, username, email, password)

	if err != nil {
		return err
	}
	return nil
}
