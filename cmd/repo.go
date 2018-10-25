package cmd

import (
	"context"
	"fmt"

	"github.com/prometheus/common/log"
	"github.com/sawadashota/gh-issue/envchain"
	"github.com/sawadashota/gh-issue/repository"
	"github.com/spf13/cobra"
)

var RepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "show writable issue repositories",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := envchain.Token()
		if err != nil {
			log.Fatal(err)
		}

		gc := repository.NewClient(context.Background(), token)
		list, err := gc.List()
		if err != nil {
			log.Fatal(err)
		}

		for _, repo := range list {
			fmt.Println(repo)
		}
	},
}
