package main

import (
	"log"
	"net/http"

	"github.com/AlexanderYAPPO/geniuslurker/geniuslurker/genius_lyrics_fetcher"
	"github.com/gorilla/mux"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/search", geniuslurker.SearchHandler).Methods("GET").Queries("q", "{q}")
	rtr.HandleFunc("/lyrics", geniuslurker.GetLyricsHandler).Methods("GET").Queries("url", "{url}")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
