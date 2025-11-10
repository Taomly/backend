package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID       int
	Username string
	Email    string
	Password string
}

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

func GetUser(db *pgxpool.Pool, username string) (*User, error) {
	user := &User{}

	err := db.QueryRow(context.Background(), `
		SELECT id, username, email, password
		FROM users
		WHERE username = $1
		`, username).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
