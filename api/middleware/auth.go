package middleware

import (
	"api-gmr/auth"
	"api-gmr/model"
	"context"
	"log"
	"net/http"
	"strings"
)

type ContextKey string

const (
	UserKey       ContextKey = "user"
	LoginPath     string     = "/login"
	AuthHeaderKey string     = "Authorization"
	BearerKey     string     = "Bearer"
)

// Auth Authenticate user with jwt-token
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == LoginPath {
			next.ServeHTTP(w, r)
			return
		}

		authHeader := r.Header.Get(AuthHeaderKey)
		if !strings.HasPrefix(authHeader, BearerKey) {
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.TrimSpace(strings.Replace(authHeader, BearerKey, "", -1))

		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			log.Println(err)
			http.Error(w, "invalid token", http.StatusBadRequest)
			return
		}

		user := model.User{}
		auth.ClaimToUser(claims, &user)

		ctx := context.WithValue(r.Context(), UserKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
