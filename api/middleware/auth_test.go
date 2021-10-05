package middleware

import (
	"api-gmr/auth"
	"api-gmr/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func authTestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func createTestToken(userId int, email, group string) string {
	user := &model.User{
		Id:    userId,
		Email: email,
		Group: group,
	}

	tokenString, err := auth.CreateToken(user)
	if err != nil {
		return ""
	}

	return tokenString
}

func TestAuth(t *testing.T) {

	type header struct {
		key   string
		value string
	}
	type args struct {
		route  string
		header header
	}

	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "login page",
			args: args{
				route:  "/v1/api/login",
				header: header{"auth", "noheader"},
			},
			want: want{http.StatusOK},
		}, {
			name: "invalid token",
			args: args{
				route:  "/v1/api/user-info",
				header: header{"Authorization", "Test"},
			},
			want: want{http.StatusBadRequest},
		}, {
			name: "valid token",
			args: args{
				route:  "/v1/api/user-info",
				header: header{"Authorization", "Bearer " + createTestToken(2, "ashadi@gmail.com", "user")},
			},
			want: want{http.StatusOK},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.args.route, nil)
			assert.NoError(t, err)
			req.Header.Set(tt.args.header.key, tt.args.header.value)

			rr := httptest.NewRecorder()

			router := mux.NewRouter()
			router.Use(Auth)
			router.HandleFunc(tt.args.route, authTestHandler)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.want.statusCode, rr.Code)
		})
	}
}
