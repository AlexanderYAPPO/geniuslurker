package geniuslurker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/AlexanderYAPPO/geniuslurker/datastructers"
)

// FetcherClient is a client for genius lurker backend
type FetcherClient struct {
	geniusLurkerFetcherHTTPClient http.Client
}

const geniusBaseURL = "https://api.genius.com/search"

var fetcherClient *FetcherClient

var onceFetcherClient sync.Once

func NewFetcherClient() *FetcherClient {
	newClient := &FetcherClient{
		geniusLurkerFetcherHTTPClient: http.Client{},
	}
	return newClient
}

// Search searches for possible songs with urls to lyrics
func (c *FetcherClient) Search(searchString string) []datastructers.SearchResult {
	httpClient := &HTTPClient{}
	req, err := http.NewRequest("GET", geniusBaseURL, nil)
	req.Header.Add("Authorization", os.Getenv("GENIUS_API_TOKEN"))
	q := req.URL.Query()
	q.Add("q", searchString)
	req.URL.RawQuery = q.Encode()
	resp, _ := httpClient.Do(req)
	var parsedJSON datastructers.BaseJSON
	err = json.NewDecoder(resp.Body).Decode(&parsedJSON)
	if err != nil {
		ErrorLogger.Println("JSON parsing error:", err)
		panic(err)
	}

	results := make([]datastructers.SearchResult, len(parsedJSON.Response.Hits), len(parsedJSON.Response.Hits))
	for index, element := range parsedJSON.Response.Hits {
		results[index] = element.Result
	}
	return results
}

// GetLyrics gets parsed lyrics for particular url
func (c *FetcherClient) GetLyrics(searchResults datastructers.SearchResult) string {
	client := &HTTPClient{}
	req, _ := http.NewRequest("GET", searchResults.URL, nil)
	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	lyrics := GetLyricsFromHTML(strings.NewReader(string(bodyBytes)))
	return searchResults.FullTitle + "\n" + lyrics
}
