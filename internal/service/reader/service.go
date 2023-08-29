package reader

import (
	"context"

	"github.com/google/uuid"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"

	"go.octolab.org/toolset/fiddle/internal/dto"
)

func New(f Fetcher) Reader {
	return reader{f}
}

type reader struct{ f Fetcher }

func (r reader) ReadArticle(ctx context.Context, url string) (*dto.Article, error) {
	resp, err := r.f.Fetch(ctx, url)
	if err != nil {
		return nil, err
	}
	defer safe.Close(resp.Body, unsafe.Ignore)

	article, err := parsers[resp.Request.URL.Hostname()].Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	article.ID = uuid.New()
	article.URL = resp.Request.URL
	article.Tags = append([]string{"reader", "article"}, article.Tags...)
	return article, nil
}
