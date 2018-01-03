package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/yappo/geniuslurker/geniuslurker"
	"log"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	res := geniuslurker.GetSearchResults(vars["q"])
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to search genius", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/search", search).Methods("GET").Queries("q", "{q}")

	http.Handle("/", rtr)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
