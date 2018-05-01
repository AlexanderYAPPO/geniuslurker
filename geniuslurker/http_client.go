package geniuslurker

import (
	"net/http"
	"strings"
)

// HTTPClient is a geniuslurker http client
type HTTPClient struct {
	http.Client
}

// Do makes request and logs
func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	// TODO: move to loggers
	InfoLogger.Println(strings.Join([]string{req.URL.String(), resp.Status, resp.Proto}, " "))
	if err != nil {
		ErrorLogger.Panicln(err)
	}
	return resp, err
}
