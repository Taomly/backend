package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StoreRefreshToken(db *pgxpool.Pool, token string, time int64) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO refresh_tokens (token, token_expires_at)
		VALUES ($1, $2)
		`, token, time)
	if err != nil {
		return err
	}
	return nil
}
