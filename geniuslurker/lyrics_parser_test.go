package geniuslurker

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestParseLyricsFromHtml(t *testing.T) {
	htmlLyrics, _ := ioutil.ReadFile("../fixtures/lyrics_1.html")
	correctParsedLyrics, _ := ioutil.ReadFile("../fixtures/lyrics_1_parsed")
	parsedLyrics := GetLyricsFromHTML(strings.NewReader(string(htmlLyrics)))
	if parsedLyrics != string(correctParsedLyrics) {
		t.Error()
	}
}
