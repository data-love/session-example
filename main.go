package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//port := os.Getenv("PORT")
	// sessionStore := sessions.NewRedisStore()
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// IndexHandler returns { "Hello": "World" }
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(`{ "Hello": "World" }`)
}
