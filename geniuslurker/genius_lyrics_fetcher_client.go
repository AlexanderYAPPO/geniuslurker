package geniuslurker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type GeniusLurkerFetcherClient struct {
	geniusLurkerFetcherHTTPClient http.Client
}

const geniusLurkerURL = "http://localhost:3000"

func NewClient() *GeniusLurkerFetcherClient {
	c := GeniusLurkerFetcherClient{
		geniusLurkerFetcherHTTPClient: http.Client{},
	}
	return &c
}

func (c *GeniusLurkerFetcherClient) Search(searchString string) []SearchResult {
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

func (c *GeniusLurkerFetcherClient) GetLyrics(searchResults SearchResult) string {
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

var geniusLurkerFetcherClient *GeniusLurkerFetcherClient

var onceGeniusLurkerFetcherClient sync.Once

// GetRedisClient returns instance of a Redis client
func GetGeniusLurkerFetcherClient() *GeniusLurkerFetcherClient {
	onceGeniusLurkerFetcherClient.Do(func() {
		geniusLurkerFetcherClient = NewClient()
	})
	return geniusLurkerFetcherClient
}
