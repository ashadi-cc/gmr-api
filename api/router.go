package api

import (
	"api-gmr/api/controller"
	"api-gmr/api/middleware"
	"api-gmr/service"
	"net/http"

	_ "api-gmr/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title GMR API
// @version 1.0
// @Description GMR API endpoint documentation

// @contact.name Ashadi
// @contact.url https://ashadi-ch.xyz
// @cotnact.email ashadi.cc@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host gmr.ashadi-ch.xyz
// @BasePath /v1/api

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
	addSwaggerRouter(r)
}

func addLoginRouter(r *mux.Router) {
	c := controller.NewLogin(service.NewAuthService())
	r.HandleFunc("/login", c.Authenticate).Methods(http.MethodPost)
}

func addUserRouter(r *mux.Router) {
	c := controller.NewUser(service.NewUserService())
	r.HandleFunc("/user-info", c.Info).Methods(http.MethodGet)
	r.HandleFunc("/user-update", c.Update).Methods(http.MethodPost)
	r.HandleFunc("/user-billing", c.Billing).Methods(http.MethodGet)
	r.HandleFunc("/user-upload", c.Upload).Methods(http.MethodPost)
}

func addSwaggerRouter(r *mux.Router) {
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
