package controller

import (
	"net/http"
	"fmt")

func Index(http.ResponseWriter, *http.Request) {
	fmt.Println("from controller")
}