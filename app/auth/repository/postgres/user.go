package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/niksche/flex/app/auth"
)

type User struct {
	ID       int    `json:"_id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r UserRepository) CreateUser(username, password string) error {

	if _, err := r.GetUser(username); err == nil {
		return fmt.Errorf("user already exists")
	}

	_, err := r.db.Exec(context.Background(), `insert into users (username, password) values ($1, $2)`, username, password)
	if err != nil {
		return fmt.Errorf("cannot insert into database")
	}
	_, err = r.db.Exec(context.Background(), `insert into accounts (username) values ($1);`, username)
	if err != nil {
		return fmt.Errorf("cannot insert into database")
	}
	return nil
}

func (r UserRepository) GetUser(username string) (auth.User, error) {
	var user auth.User
	err := r.db.QueryRow(context.Background(),
		`	SELECT username, password
				FROM users
				WHERE username = $1`, username,
	).Scan(
		&user.Username,
		&user.Password,
	)

	if err != nil {
		return user, fmt.Errorf("cannot find that user")
	}
	return user, nil
}

func (r UserRepository) MakePayment(username string) error {

	if _, err := r.GetUser(username); err != nil {
		return fmt.Errorf("cannot find person with thhat username")
	}

	_, err := r.db.Exec(context.Background(), `UPDATE accounts SET balance = balance - round(1.1 , 1) WHERE username = $1`, username)
	if err != nil {
		return fmt.Errorf("cannot spend that much money")
	}

	return nil
}
