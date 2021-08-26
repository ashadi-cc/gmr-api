package api

import (
	"api-gmr/env"
	"fmt"
	"log"
	"net/http"
)

// Run api
func Run() {
	r := newRouter()

	port := env.GetValue("APP_PORT", "8080")
	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
