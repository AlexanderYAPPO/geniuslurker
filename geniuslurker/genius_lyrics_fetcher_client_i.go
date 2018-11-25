package geniuslurker

import "github.com/AlexanderYAPPO/geniuslurker/datastructers"

type FetcherClientI interface {
	Search(searchString string) []datastructers.SearchResult
	GetLyrics(searchResults datastructers.SearchResult) string
}

// GetFetcherClient returns instance of a Fetcher client
func GetFetcherClient() FetcherClientI {
	onceFetcherClient.Do(func() {
		fetcherClient = NewFetcherClient()
	})
	return fetcherClient
}
