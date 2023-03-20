package middlewares

import "net/http"

type contextKey int
const authenticatedUserKey contextKey = 0

type EnsureAuth struct {
	handler http.Handler
}

func (ea *EnsureAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
}