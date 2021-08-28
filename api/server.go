package api

import (
	"api-gmr/config"
	"fmt"
	"log"
	"net/http"
)

//Run run api server
func Run() {
	r := newRouter()

	port := config.GetApp().AppPort
	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
