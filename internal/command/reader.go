package command

import (
	"net/http"

	"github.com/spf13/cobra"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"

	"go.octolab.org/toolset/fiddle/internal/pkg/cookiejar"
	"go.octolab.org/toolset/fiddle/internal/service/fetcher"
	"go.octolab.org/toolset/fiddle/internal/service/presenter"
	"go.octolab.org/toolset/fiddle/internal/service/reader"
)

func Read() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "read",
		Args: cobra.NoArgs,
	}

	var glamoured bool
	article := &cobra.Command{
		Use:  "article",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			j := cookiejar.Must(
				new(cookiejar.Options).StoreToDisk().AsFile("bin/cookies.json"),
			)
			c := &http.Client{Jar: j}
			r := reader.New(fetcher.New(c))
			defer safe.Do(j.Dump, unsafe.Ignore)

			article, err := r.ReadArticle(cmd.Context(), args[0])
			if err != nil {
				return err
			}

			var opts []presenter.Option
			if glamoured {
				opts = append(opts, presenter.Glamoured())
			}
			return presenter.New(cmd.OutOrStdout(), opts...).Article(*article)
		},
	}
	fs := article.Flags()
	fs.BoolVar(&glamoured, "glamoured", false, "use it if you prefer to read article in terminal")

	cmd.AddCommand(article)

	return cmd
}
