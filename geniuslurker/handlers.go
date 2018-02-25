package geniuslurker

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func fetchLyricsFromGenius(url string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	lyrics := GetLyricsFromHTML(strings.NewReader(string(bodyBytes)))
	return lyrics
}

func GetLyricsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url := vars["url"]
	lyrics := fetchLyricsFromGenius(url)
	io.WriteString(w, lyrics)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Failed to get search results:", r)
			http.Error(w, "Failed to search Genius", http.StatusInternalServerError)
			return
		}
	}()
	vars := mux.Vars(r)
	res := GetSearchResults(vars["q"])
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to search genius", http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
