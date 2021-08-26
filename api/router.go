package api

import (
	"api-gmr/api/controller"
	"api-gmr/api/middleware"
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
	c := controller.NewLogin()
	r.HandleFunc("/login", c.PostLogin).Methods(http.MethodPost)
}
