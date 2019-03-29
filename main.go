package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ResponseFormat struct {
	Data       []User `json:"data,omitempty"`
	RunTime    string `json:"run_time,omitempty"`
	HttpMethod string `json:"http_method,omitempty"`
	HttpPath   string `json:"http_path,omitempty"`
}

type User struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}

var users []User

func listUsers(w http.ResponseWriter, r *http.Request) {

	var response = new(ResponseFormat)

	t := time.Now()
	response.RunTime = t.String()

	response.Data = append(response.Data, User{Username: "afhdfgbdf", Email: "adfbdfbdfbd@a.a"})
	response.Data = append(response.Data, User{Username: "bdfbdfbdf", Email: "bfbdfbdfbdfb@b.b"})
	response.Data = append(response.Data, User{Username: "cbdfbdfbfbdfb"})

	response.HttpMethod = r.Method
	url := r.URL
	response.HttpPath = url.String()

	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", listUsers)
	log.Fatal(http.ListenAndServe(":12345", router))
}
