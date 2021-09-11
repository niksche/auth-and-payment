package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/niksche/flex/app/auth"
)

type PaymentRepository struct {
	db *pgxpool.Pool
}

func NewPaymentRepository(db *pgxpool.Pool) *PaymentRepository {
	return &PaymentRepository{
		db: db,
	}
}
func (r PaymentRepository) MakePayment(username string) error {

	account, err := r.GetAccount(username)
	if err != nil {
		return fmt.Errorf("cannot find person with thhat username")
	}

	if account.AccountBalance < 1.1 {
		return fmt.Errorf("not enough money")
	}

	_, err = r.db.Exec(context.Background(), `UPDATE accounts SET balance = balance - round(1.1 , 1) WHERE username = $1`, username)
	if err != nil {
		return fmt.Errorf("cannot spend that much money")
	}

	return nil
}

func (r PaymentRepository) GetAccount(username string) (auth.User, error) {
	var user auth.User
	err := r.db.QueryRow(context.Background(),
		`	SELECT username, balance
				FROM accounts
				WHERE username = $1`, username,
	).Scan(
		&user.Username,
		&user.AccountBalance,
	)

	if err != nil {
		return user, fmt.Errorf("cannot find that account")
	}
	return user, nil
}
