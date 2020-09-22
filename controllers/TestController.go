package controllers

import (
	"fmt"
	"net/http"

	"github.com/rohimihsan/auth-sys/config/db"
)

func TestUp(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Server is up")
}

func TestConnection(w http.ResponseWriter, r *http.Request) {
	_, err := db.Con()

	if err != nil {
		fmt.Fprint(w, "Error connecting to db")
		return
	}

	fmt.Fprint(w, "Connection success")
	return
}
