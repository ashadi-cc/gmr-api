package api

import (
	"api-gmr/api/controller"
	"api-gmr/api/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	HandleFunc(w http.ResponseWriter, r *http.Request)
}

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
	addRouter(r, controller.NewLogin(), "/login", http.MethodPost)
}

func addRouter(r *mux.Router, c Controller, path string, methods ...string) {
	r.HandleFunc(path, c.HandleFunc).Methods(methods...)
}
