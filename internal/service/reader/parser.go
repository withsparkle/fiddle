package reader

import (
	"io"

	"go.octolab.org/toolset/fiddle/internal/dto"
	"go.octolab.org/toolset/fiddle/internal/service/reader/strategy"
)

type parser func(io.Reader) (*dto.Article, error)

func (fn parser) Parse(body io.Reader) (*dto.Article, error) { return fn(body) }

var parsers = map[string]parser{
	"":               strategy.Default,
	"fastfounder.ru": strategy.FastFounder,
}
