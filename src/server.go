package src

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")

	log.Println("[Server] HTTP is served in :" + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
