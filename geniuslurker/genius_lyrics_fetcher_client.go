package geniuslurker

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

// FetcherClient is a client for genius lurker backend
type FetcherClient struct {
	geniusLurkerFetcherHTTPClient http.Client
}

const geniusBaseURL = "https://api.genius.com/search"
const geniusToken = "Bearer us4hrg63-ZYFCFmecW9iS3nXoLs5rkTkFIGhECwNHtMda0GyCINDkleGdmiKjAmx"

var fetcherClient *FetcherClient

var onceFetcherClient sync.Once

// GetFetcherClient returns instance of a Genius Lyrics Fetcher client
func GetFetcherClient() *FetcherClient {
	onceFetcherClient.Do(func() {
		fetcherClient = &FetcherClient{
			geniusLurkerFetcherHTTPClient: http.Client{},
		}
	})
	return fetcherClient
}

// Search searches for possible songs with urls to lyrics
func (c *FetcherClient) Search(searchString string) []SearchResult {
	httpClient := &HTTPClient{}
	req, err := http.NewRequest("GET", geniusBaseURL, nil)
	req.Header.Add("Authorization", geniusToken)
	q := req.URL.Query()
	q.Add("q", searchString)
	req.URL.RawQuery = q.Encode()
	resp, _ := httpClient.Do(req)
	var parsedJSON baseJSON
	err = json.NewDecoder(resp.Body).Decode(&parsedJSON)
	if err != nil {
		ErrorLogger.Println("JSON parsing error:", err)
		panic(err)
	}

	results := make([]SearchResult, len(parsedJSON.Response.Hits), len(parsedJSON.Response.Hits))
	for index, element := range parsedJSON.Response.Hits {
		results[index] = element.Result
	}
	return results
}

// GetLyrics gets parsed layrics for particular url
func (c *FetcherClient) GetLyrics(searchResults SearchResult) string {
	client := &HTTPClient{}
	req, _ := http.NewRequest("GET", searchResults.URL, nil)
	resp, _ := client.Do(req)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	lyrics := GetLyricsFromHTML(strings.NewReader(string(bodyBytes)))
	return searchResults.FullTitle + "\n" + lyrics
}

// SearchResult represents search result from genius
type SearchResult struct {
	FullTitle string `json:"full_title"`
	URL       string `json:"url"`
}

type hitJSON struct {
	Result SearchResult `json:"result"`
}

type responseJSON struct {
	Hits []hitJSON `json:"hits"`
}

type baseJSON struct {
	Response responseJSON `json:"response"`
}
