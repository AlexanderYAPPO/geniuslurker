package datastructers

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

type BaseJSON struct {
	Response responseJSON `json:"response"`
}
