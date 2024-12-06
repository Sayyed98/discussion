package main

import (
	"discussion/auth"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/signup", auth.Signup).Methods("GET")
	fmt.Println("server listening on port 8080")
	http.ListenAndServe(":8080", route)
}
