package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.Handle("/", r)
	fmt.Println("Listening on port 8080", "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
