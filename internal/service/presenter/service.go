package presenter

import (
	"io"
	"text/template"

	converter "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/charmbracelet/glamour"

	"go.octolab.org/toolset/fiddle/internal/dto"
)

func New(out io.Writer, opts ...Option) Presenter {
	markdown := converter.NewConverter("", true, nil)
	sparkle := template.New("sparkle")
	sparkle.Funcs(map[string]interface{}{
		"markdownify": markdown.ConvertString,
	})
	template.Must(sparkle.ParseFS(fs, "templates/sparkle/*.gomd"))

	var opt option
	for _, config := range opts {
		opt = config(opt)
	}
	return Presenter{out, sparkle, opt}
}

type Presenter struct {
	o io.Writer
	r Renderer
	c option
}

func (p Presenter) Article(article dto.Article) error {
	if !p.c.glamoured {
		return p.r.ExecuteTemplate(p.o, "article.gomd", article)
	}

	r, err := glamour.NewTermRenderer(glamour.WithStyles(glamour.DraculaStyleConfig))
	if err != nil {
		return err
	}
	if err := p.r.ExecuteTemplate(r, "article.gomd", article); err != nil {
		return err
	}
	if err := r.Close(); err != nil {
		return nil
	}
	_, err = io.Copy(p.o, r)
	return err
}
