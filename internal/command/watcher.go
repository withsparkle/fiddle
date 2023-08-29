package command

import "github.com/spf13/cobra"

func Watch() *cobra.Command {
	watcher := &cobra.Command{
		Use:  "watch",
		Args: cobra.NoArgs,
	}

	watcher.AddCommand(&cobra.Command{
		Use:  "youtube",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("watch youtube")
			return nil
		},
	})

	return watcher
}
