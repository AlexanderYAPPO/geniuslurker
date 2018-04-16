package geniuslyricsfetcher

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/AlexanderYAPPO/geniuslurker/geniuslurker"
	"github.com/gorilla/mux"
)

func fetchLyricsFromGenius(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	//TODO: move to loggers
	geniuslurker.InfoLogger.Println(strings.Join([]string{req.URL.String(), resp.Status, resp.Proto}, " "))
	if err != nil {
		geniuslurker.ErrorLogger.Println(err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	lyrics := GetLyricsFromHTML(strings.NewReader(string(bodyBytes)))
	return lyrics
}

//GetLyricsHandler is a handler for /lyrics
func GetLyricsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	lyrics := fetchLyricsFromGenius(url)
	io.WriteString(w, lyrics)
}

//SearchHandler is a handler for /search
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			geniuslurker.ErrorLogger.Println("Failed to get search results:", r)
			http.Error(w, "Failed to search Genius", http.StatusInternalServerError)
			return
		}
	}()
	vars := mux.Vars(r)
	res := GetSearchResults(vars["q"])
	b, err := json.Marshal(res)
	if err != nil {
		geniuslurker.ErrorLogger.Println(err)
		http.Error(w, "Failed to search genius", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
