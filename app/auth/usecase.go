package auth

type UseCase interface {
	SignUp(username, password string) error
	LoginUser(username, password string) error
}
