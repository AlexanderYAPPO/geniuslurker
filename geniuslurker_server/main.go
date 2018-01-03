package main

import (
	"github.com/gorilla/mux"
	"github.com/yappo/geniuslurker/geniuslurker"
	"log"
	"net/http"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/search", geniuslurker.SearchHandler).Methods("GET").Queries("q", "{q}")
	rtr.HandleFunc("/lyrics", geniuslurker.GetLyricsHandler).Methods("GET").Queries("url", "{url}")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
