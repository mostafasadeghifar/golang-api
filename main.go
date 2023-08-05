package main

import (
	"log"
	"myjoke/config"
	"myjoke/route"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	// route.Routing(router)
	route.JokeRout(router)
	config.Connect();

	log.Fatal(http.ListenAndServe(":8083",router))
}
