package strategy

import (
	"io"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"

	"go.octolab.org/toolset/fiddle/internal/dto"
)

func Default(body io.Reader) (*dto.Article, error) {
	root, err := html.Parse(body)
	if err != nil {
		return nil, err
	}
	doc := goquery.NewDocumentFromNode(root)

	article := new(dto.Article)
	article.Title = doc.Find("title").Text()
	article.Content, err = doc.Find("body").Html()
	return article, err
}
