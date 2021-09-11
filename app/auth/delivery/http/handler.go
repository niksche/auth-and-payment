package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/niksche/flex/app/auth"
)

var (
	jwtKey = []byte("secretkeydonttellanyonepls")
)

type Handler struct {
	useCase auth.UseCase
}

type Claims struct {
	Username           string `json:"username"`
	jwt.StandardClaims ``
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var userData auth.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := h.useCase.SignUp(userData.Username, userData.Password); err != nil {
		// if err := h.useCase.SignUp("ef", "fe"); err != nil {
		w.Write([]byte("error while creating new profile"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("succesfully created"))
}

func (h Handler) LogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}

	var userData auth.User
	err := json.NewDecoder(r.Body).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := h.useCase.LoginUser(userData.Username, userData.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Username: userData.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})

	w.Write([]byte(tokenString))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("succesfully loged in"))

}

func (h Handler) LogOut(w http.ResponseWriter, r *http.Request) {
	expirationTime := time.Now().Add(-time.Hour)
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   "",
			Expires: expirationTime,
		})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("succesfully logged out"))
}
