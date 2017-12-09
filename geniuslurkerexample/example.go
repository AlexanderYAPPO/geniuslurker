package main

import (
	"fmt"
	"github.com/yappo/geniuslurker/geniuslurker"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	res := geniuslurker.GetSearchResults("Madvillain")
	fmt.Println(res[0].Url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", res[0].Url, nil)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("whoops:", err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	lyrics := geniuslurker.GetLyricsFromHTML(strings.NewReader(string(bodyBytes)))
	fmt.Println(lyrics)
}
