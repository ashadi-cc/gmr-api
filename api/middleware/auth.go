package middleware

import (
	"api-gmr/auth"
	"api-gmr/model"
	"context"
	"log"
	"net/http"
	"strings"
)

//ContextKey represents context data type
type ContextKey string

const (
	//UserKey represents user context key
	UserKey ContextKey = "user"
	//LoginPath represents login route path
	LoginPath string = "/login"
	//AuthHeaderKey represents Authorization header key value
	AuthHeaderKey string = "Authorization"
	//BearerKey represents token header key method
	BearerKey string = "Bearer"
)

// Auth represents middleware with validate jwt-token.
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
