package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func StoreRefreshToken(db *pgxpool.Pool, userId int, token string, time int64) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO refresh_tokens (user_id, token, token_expires_at)
		VALUES ($1, $2, $3)
		`, userId, token, time)
	if err != nil {
		return err
	}
	return nil
}

func RestoreRefreshToken(db *pgxpool.Pool, oldToken, newToken string) error {
	_, err := db.Exec(context.Background(), `
		UPDATE refresh_tokens
		SET token = $1
		WHERE token = $2
`, newToken, oldToken)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllUserTokens(db *pgxpool.Pool, userId int) error {
	_, err := db.Exec(context.Background(), "DELETE FROM refresh_tokens WHERE user_id = $1", userId)
	if err != nil {
		return err
	}
	return nil
}
