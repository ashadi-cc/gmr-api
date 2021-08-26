package api

import (
	"api-gmr/api/controller"
	"api-gmr/api/middleware"
	"api-gmr/service"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	addMiddleware(r)
	addRouters(r)
	return r
}

func addMiddleware(r *mux.Router) {
	r.Use(
		middleware.Auth,
	)
}

func addRouters(r *mux.Router) {
	addLoginRouter(r)
}

func addLoginRouter(r *mux.Router) {
	c := controller.NewLogin(service.NewAuthService())
	r.HandleFunc("/login", c.Authenticate).Methods(http.MethodPost)
}
