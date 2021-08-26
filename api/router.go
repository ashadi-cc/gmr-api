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
	addUserRouter(r)
}

func addLoginRouter(r *mux.Router) {
	c := controller.NewLogin(service.NewAuthService())
	r.HandleFunc("/login", c.Authenticate).Methods(http.MethodPost)
}

func addUserRouter(r *mux.Router) {
	c := controller.NewUser(service.NewUserService())
	r.HandleFunc("/user-info", c.Info).Methods(http.MethodGet)
	r.HandleFunc("/user-update", c.Update).Methods(http.MethodPost)
}
