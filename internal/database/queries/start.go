package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTable(db *pgxpool.Pool) error {
	_, err := db.Exec(context.Background(), `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT UNIQUE NOT NULL,
		email TEXT UNIQUE NOT NULL,
		password TEXT NOT NULL
	);
	`)
	if err != nil {
		return err
	}
	return nil
}
