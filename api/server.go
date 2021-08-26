package api

import (
	"fmt"
	"log"
	"net/http"
)

var port = "8080"

// Run api
func Run() {
	r := newRouter()

	log.Println("Listening on port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
