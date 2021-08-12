package src

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deta-aditya/lunaticup-be/src/tournaments"
	"github.com/gorilla/mux"
)

type Tournament struct {
	Name    string   `json:"name"`
	Method  string   `json:"method"`
	Players []string `json:"players"`
}

func RunServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Lunaticup API")
	})

	r.HandleFunc("/tournaments", tournaments.HandleCreate).Methods("POST")

	log.Println("[Server] HTTP is served in :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
