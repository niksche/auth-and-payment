package auth

type UseCase interface {
	SignUp(username, password string) error
	LoginUser(username, password string) error
	MakePayment(username string) error
}
