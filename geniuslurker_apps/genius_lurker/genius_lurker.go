package main

import (
	"net/http"

	"github.com/AlexanderYAPPO/geniuslurker/geniuslurker"
	"github.com/AlexanderYAPPO/geniuslurker/geniuslurker/geniuslyricsfetcher"
	"github.com/gorilla/mux"
)

func main() {
	geniuslurker.InitLoggers()
	rtr := mux.NewRouter()
	rtr.HandleFunc("/search", geniuslyricsfetcher.SearchHandler).Methods("GET").Queries("q", "{q}")
	rtr.HandleFunc("/lyrics", geniuslyricsfetcher.GetLyricsHandler).Methods("GET").Queries("url", "{url}")

	http.Handle("/", rtr)

	http.ListenAndServe(":3000", nil)
}
