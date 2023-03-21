package middlewares

import (
	"context"
	"net/http"
	"os"
	. "server/types"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey int
const authenticatedUserKey contextKey = 0

type EnsureAuth struct {
	handler http.Handler
}

func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := c.Value

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return os.Getenv("JWT_SECRET"), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var user = map[string]string{
		"user1": "password1",
	}

	newContext := context.WithValue(r.Context(), authenticatedUserKey, user)
	newRequest := r.WithContext(newContext)

	ea.handler.ServeHTTP(w, newRequest)
}

func Protect(handler http.Handler) *EnsureAuth {
	return &EnsureAuth{handler}
}