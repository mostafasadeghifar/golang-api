package main

import (
	"fmt"
	"log"
	// "myjoke/route"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();

	router.HandleFunc("/tt",TestFunc)
	log.Fatal(http.ListenAndServe(":8083",router))
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hiii")
	fmt.Println("hiii")
}