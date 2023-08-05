package route

import (
	"fmt"
	"myjoke/controller"
	"net/http"

	"github.com/gorilla/mux"
)
var JokeRout = func(r *mux.Router){
	r.HandleFunc("/samples", SampleHandler)
	r.HandleFunc("/index", controller.Index)
	r.HandleFunc("/create", controller.Create).Methods("POST")
}
// func Routing(router *mux.Router){
// 	router.HandleFunc("/samples", SampleHandler)
// 	router.HandleFunc("/index", controller.Index)
// }

func SampleHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"hi babe")
}

func Sample() {
	fmt.Println("Hello form router!")
}