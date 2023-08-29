package strategy

import (
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/google/uuid"

	"go.octolab.org/toolset/fiddle/internal/dto"
)

func FastFounder(body io.Reader) (*dto.Article, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return nil, err
	}
	if doc.Find("#comments").Length() == 0 {
		return nil, fmt.Errorf("unauthorized. please update credentials")
	}

	article := new(dto.Article)
	article.Date, err = time.Parse("02.01.2006", doc.Find("div.post_date").Text())
	if err != nil {
		return nil, err
	}
	doc.Find(`div.tags_link a[rel="tag"]`).Each(func(_ int, node *goquery.Selection) {
		raw, present := node.Attr("href")
		if !present {
			return
		}
		u, err := url.Parse(raw)
		if err != nil {
			return
		}
		article.Links = append(article.Links, dto.Link{
			ID:   uuid.New(),
			URL:  u,
			Type: dto.Taxonomy,
			Name: strings.Trim(node.Text(), "#"),
		})
	})
	article.Title = doc.Find("h1.entry-title").Text()
	content := doc.
		Find("div.after-entry").
		NextUntil("p:last-of-type").
		WrapAllHtml("<div></div>").Parent()
	content.Find("a[href]").Each(func(_ int, node *goquery.Selection) {
		raw, present := node.Attr("href")
		if !present {
			return
		}
		u, err := url.Parse(raw)
		if err != nil {
			return
		}
		article.Links = append(article.Links, dto.Link{
			ID:   uuid.New(),
			URL:  u,
			Type: dto.Reference,
			Name: node.Text(),
		})
	})
	article.Content, err = content.Html()
	article.Summary = "Coming soon ⏳"
	return article, err
}
