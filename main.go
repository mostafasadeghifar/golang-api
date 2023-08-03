package main

import (
	"fmt"
	"log"
	"myjoke/config"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();
	config.Connect();

	router.HandleFunc("/tt",TestFunc)
	log.Fatal(http.ListenAndServe(":8083",router))
}

func TestFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("hiii")
	fmt.Fprint(w,"lkjalsdf")
	fmt.Println("hiii")
}