package main

import (
	"log"
	"myjoke/route"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// route.Routing(router)
	route.JokeRout(router)
	log.Fatal(http.ListenAndServe(":8083",router))
}
