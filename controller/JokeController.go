package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"myjoke/model"
)

func Index(http.ResponseWriter, *http.Request) {
	fmt.Println("from controller")
}

func Create(w http.ResponseWriter,req *http.Request) {
	CreateJoke := &model.Joke{}
	ParseBody(req, CreateJoke)
	fmt.Println(&CreateJoke)
	b := CreateJoke.CreateJoke()
	res, _ := json.Marshal(b)
	w.Write(res)

}

func ParseBody(r *http.Request,X interface{}){
	if body, err := ioutil.ReadAll(r.Body); err==nil{
		if err:= json.Unmarshal([]byte(body),X); err != nil {
			return 
		}
	}
}
