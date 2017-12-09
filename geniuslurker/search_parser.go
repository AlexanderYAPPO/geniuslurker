package geniuslurker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO: change it to some url type
const geniusBaseURL string = "https://api.genius.com/search?q="

type ResultJSON struct {
	FullTitle string `json:"full_title"`
	Url       string `json:"url"`
}

type HitJSON struct {
	Result ResultJSON `json:"result"`
}

type ResponseJSON struct {
	Hits []HitJSON `json:"hits"`
}

type BaseJSON struct {
	Response ResponseJSON `json:"response"`
}

// Returns search results
func GetSearchResults(searchString string) []ResultJSON {
	var tmpBuffer bytes.Buffer
	tmpBuffer.WriteString(geniusBaseURL)
	tmpBuffer.WriteString(searchString)
	geniusSearchURL := tmpBuffer.String()

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", geniusSearchURL, nil)
	req.Header.Add("Authorization", "Bearer us4hrg63-ZYFCFmecW9iS3nXoLs5rkTkFIGhECwNHtMda0GyCINDkleGdmiKjAmx")
	resp, err := httpClient.Do(req)
	searchBody := resp.Body
	defer searchBody.Close()

	tmpBufferP := new(bytes.Buffer)
	tmpBufferP.ReadFrom(searchBody)
	jsonArray := tmpBufferP.Bytes()

	var parsedJSON = new(BaseJSON)
	err = json.Unmarshal(jsonArray, &parsedJSON)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	results := make([]ResultJSON, len(parsedJSON.Response.Hits), len(parsedJSON.Response.Hits))
	for index, element := range parsedJSON.Response.Hits {
		results[index] = element.Result
	}
	return results
}
