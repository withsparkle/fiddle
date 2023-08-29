package reader

import (
	"context"
	"io"
	"net/http"

	"go.octolab.org/toolset/fiddle/internal/dto"
)

type Fetcher interface {
	Fetch(context.Context, string) (*http.Response, error)
}

type Parser interface {
	Parse(io.Reader) (*dto.Article, error)
}

type Reader interface {
	ReadArticle(context.Context, string) (*dto.Article, error)
}
