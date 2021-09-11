package auth

type User struct {
	ID             int     `json:"_id,omitempty"`
	Username       string  `json:"username"`
	Password       string  `json:"password"`
	AccountBalance float32 `json:"balance,omitempty"`
}
type UserRepository interface {
	CreateUser(username, password string) error
	GetUser(username string) (User, error)
}
