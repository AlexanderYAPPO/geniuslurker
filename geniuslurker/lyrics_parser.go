package geniuslurker

import (
	"bytes"
	"golang.org/x/net/html"
	"io"
)

func getLyricsBlockFlag(currentToken html.TokenType) bool {
	switch currentToken {
	case html.StartTagToken:
		return true
	case html.EndTagToken:
		return false
	}
	return false
}

func checkIsStringEmpty(s string) bool {
	if s == "\n" {
		return false
	}
	for _, v := range s {
		if v != 32 && v != 10 {
			return false
		}
	}
	return true
}

func checkIfLyricsBlock(pageTokenizer *html.Tokenizer) bool {
	tagName, hasAttrs := pageTokenizer.TagName()
	if string(tagName) != "div" || !hasAttrs {
		return false
	}

	for tagName, tagValue, isMore := pageTokenizer.TagAttr(); ; tagName, tagValue, isMore = pageTokenizer.TagAttr() {
		if string(tagName) == "class" && string(tagValue) == "lyrics" {
			return true
		}
		if !isMore {
			break
		}
	}
	return false
}

// Parse lyrics from Genius HTML page and returns the as string (including \n and other stuff)
func GetLyricsFromHTML(htmlDataReader io.Reader) string {
	pageTokenizer := html.NewTokenizer(htmlDataReader)
	var resultBuffer bytes.Buffer
	lyricsBlockFlag := false

	for {
		currentTag := pageTokenizer.Next()

		switch currentTag {
		case html.ErrorToken:
			if pageTokenizer.Err() == io.EOF {
				return resultBuffer.String()
			}
			// Unexpected case
			panic(pageTokenizer.Err())
		}

		if !lyricsBlockFlag && checkIfLyricsBlock(pageTokenizer) {
			lyricsBlockFlag = getLyricsBlockFlag(currentTag)
		}
		// After this we've already handled and read the tag

		if !lyricsBlockFlag {
			continue
		}

		// Handling the next tag
		tagName, _ := pageTokenizer.TagName()

		// Checking if the lyrics block ends
		if string(tagName) == "div" {
			switch currentTag {
			case html.EndTagToken:
				lyricsBlockFlag = false
				continue
			}
		}

		switch currentTag {
		case html.TextToken:
			// The only type of tags we're interested in
			currentToken := pageTokenizer.Token()
			lyricsLine := html.UnescapeString(currentToken.String())
			if checkIsStringEmpty(lyricsLine) {
				continue
			}
			resultBuffer.WriteString(lyricsLine)
		}
	}

}
