package geniuslurker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

// FetcherClient is a client for genius lurker backend
type FetcherClient struct {
	geniusLurkerFetcherHTTPClient http.Client
}

const geniusLurkerURL = "http://localhost:3000"

// NewClient creates a new client
func NewClient() *FetcherClient {
	c := FetcherClient{
		geniusLurkerFetcherHTTPClient: http.Client{},
	}
	return &c
}

// Search searches for possible songs
func (c *FetcherClient) Search(searchString string) []SearchResult {
	searchURL, _ := url.Parse(geniusLurkerURL + "/search")
	query := searchURL.Query()

	query.Set("q", searchString)
	searchURL.RawQuery = query.Encode()
	req, _ := http.NewRequest("GET", searchURL.String(), nil)

	resp, err := c.geniusLurkerFetcherHTTPClient.Do(req)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	defer resp.Body.Close()
	var searchResults []SearchResult
	json.NewDecoder(resp.Body).Decode(&searchResults)
	fmt.Println(searchResults, err)
	return searchResults
}

// GetLyrics gets parsed layrics for particular url
func (c *FetcherClient) GetLyrics(searchResults SearchResult) string {
	searchURL, _ := url.Parse(geniusLurkerURL + "/lyrics")
	query := searchURL.Query()

	query.Set("url", searchResults.URL)
	searchURL.RawQuery = query.Encode()
	req, _ := http.NewRequest("GET", searchURL.String(), nil)

	resp, err := c.geniusLurkerFetcherHTTPClient.Do(req)
	if err != nil {
		fmt.Println("whoops:", err)
		panic(err)
	}
	defer resp.Body.Close()
	var lyrics string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		lyrics = string(bodyBytes)
	}
	return searchResults.FullTitle + "\n" + lyrics
}

// SearchResult represents search result from genius
type SearchResult struct {
	FullTitle string `json:"full_title"`
	URL       string `json:"url"`
}

var fetcherClient *FetcherClient

var onceFetcherClient sync.Once

// GetFetcherClient returns instance of a Genius Lyrics Fetcher client
func GetFetcherClient() *FetcherClient {
	onceFetcherClient.Do(func() {
		fetcherClient = NewClient()
	})
	return fetcherClient
}
