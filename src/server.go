package src

import (
	"fmt"
	"log"
	"net/http"

	"example.com/lunaticup-be/src/tournaments"
	"github.com/gorilla/mux"
)

type Tournament struct {
	Name    string   `json:"name"`
	Method  string   `json:"method"`
	Players []string `json:"players"`
}

func Run() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Lunaticup API")
	})

	r.HandleFunc("/tournaments", tournaments.HandleCreate).Methods("POST")

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(rw, "Lunaticup API")
	// })

	// http.HandleFunc("/tournaments", func(rw http.ResponseWriter, r *http.Request) {
	// 	tournamentsJSON := []Tournament{
	// 		{Name: "The Slumber Cup", Method: "Single Elimination", Players: []string{"BunnySlice", "StupidFern", "Imbechole"}},
	// 		{Name: "Naruto Summer Cup", Method: "Single Elimination", Players: []string{"Benny", "Bensos", "Brian", "Baba"}},
	// 	}

	// 	rw.Header().Set("Content-Type", "application/json")
	// 	json.NewEncoder(rw).Encode(tournamentsJSON)
	// })

	log.Fatal(http.ListenAndServe(":8080", r))
}
