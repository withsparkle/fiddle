package fetcher

import "net/http"

type Client interface {
	Do(*http.Request) (*http.Response, error)
}

type Validator interface {
	Validate(*http.Response) error
}
