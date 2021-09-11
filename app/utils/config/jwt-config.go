package config

import "github.com/dgrijalva/jwt-go"

var (
	JwtKey = []byte("secretkeydonttellanyonepls")
)

type Claims struct {
	Username           string `json:"username"`
	jwt.StandardClaims ``
}
