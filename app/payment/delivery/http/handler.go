package http

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/niksche/flex/app/payment"
	jwtConfig "github.com/niksche/flex/app/utils/config"
)

type Handler struct {
	useCase payment.UseCase
}

func NewHandler(useCase payment.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h Handler) Payment(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &jwtConfig.Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtConfig.JwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err := h.useCase.MakePayment(claims.Username); err != nil {
		w.WriteHeader(http.StatusPaymentRequired)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s\nMoney's spent", claims.Username)))

}
