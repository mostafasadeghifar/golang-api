package route

import (
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
)

func Routing(router *mux.Router){
	router.HandleFunc("/samples", SampleHandler)
}

func SampleHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"hi babe")
}

func Sample() {
	fmt.Println("Hello form router!")
}