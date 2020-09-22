package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohimihsan/go-auth-jwt/controllers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.TestUp).Methods("GET")

	test := r.PathPrefix("/test").Subrouter()
	test.HandleFunc("/", controllers.TestUp).Methods("GET")
	test.HandleFunc("/conn", controllers.TestConnection).Methods("GET")

	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/home", controllers.Home).Methods("GET")
	r.HandleFunc("/refresh", controllers.RefreshToken).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}
