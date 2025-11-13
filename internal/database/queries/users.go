package queries

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
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
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return fmt.Errorf("user with this username or email already exists")
			}
		}
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

func GetUserByID(db *pgxpool.Pool, id int) (*User, error) {
	user := &User{}

	err := db.QueryRow(context.Background(), `
		SELECT id, username, email, password
		FROM users
		WHERE id = $1
		`, id).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUserPassword(db *pgxpool.Pool, id int, password string) error {
	cmdTag, err := db.Exec(context.Background(), `
		UPDATE users
		SET password = $1
		WHERE id = $2
		`, password, id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() != 1 {
		return errors.New("couldn't update password")
	}

	return nil
}
