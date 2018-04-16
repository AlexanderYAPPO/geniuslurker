package geniuslyricsfetcher

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/AlexanderYAPPO/geniuslurker/geniuslurker"
)

const geniusBaseURL = "https://api.genius.com/search"
const geniusToken = "Bearer us4hrg63-ZYFCFmecW9iS3nXoLs5rkTkFIGhECwNHtMda0GyCINDkleGdmiKjAmx"

// ResultJSON represents an element of a found song
type ResultJSON struct {
	FullTitle string `json:"full_title"`
	Url       string `json:"url"`
}

type hitJSON struct {
	Result ResultJSON `json:"result"`
}

type responseJSON struct {
	Hits []hitJSON `json:"hits"`
}

type baseJSON struct {
	Response responseJSON `json:"response"`
}

// GetSearchResults requests Genius API with a search request
func GetSearchResults(searchString string) []ResultJSON {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", geniusBaseURL, nil)
	req.Header.Add("Authorization", geniusToken)
	q := req.URL.Query()
	q.Add("q", searchString)
	req.URL.RawQuery = q.Encode()
	geniuslurker.InfoLogger.Println("Search Url: ", req.URL.String())
	resp, err := httpClient.Do(req)
	//TODO: move to loggers
	geniuslurker.InfoLogger.Println(strings.Join([]string{req.URL.String(), resp.Status, resp.Proto}, " "))
	if err != nil {
		geniuslurker.ErrorLogger.Println("Request error: ", err)
		panic(err)
	}

	var parsedJSON baseJSON
	err = json.NewDecoder(resp.Body).Decode(&parsedJSON)
	if err != nil {
		geniuslurker.ErrorLogger.Println("JSON parsing error:", err)
		panic(err)
	}

	results := make([]ResultJSON, len(parsedJSON.Response.Hits), len(parsedJSON.Response.Hits))
	for index, element := range parsedJSON.Response.Hits {
		results[index] = element.Result
	}
	return results
}
