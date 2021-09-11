package payment

import "github.com/niksche/flex/app/auth"

type UserRepository interface {
	MakePayment(username string) error
	GetAccount(username string) (auth.User, error)
}
