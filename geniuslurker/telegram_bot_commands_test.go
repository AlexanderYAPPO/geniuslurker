package geniuslurker

import (
	"io/ioutil"
	"testing"
)

func TestSplitTextOnBlocks(t *testing.T) {
	lyrics, _ := ioutil.ReadFile("../fixtures/long_lyrics")
	blocks := splitTextOnBlocks(string(lyrics))
	if len(blocks) != 2 {
		t.Error()
	}
}