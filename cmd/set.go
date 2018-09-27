package cmd

import (
	"os"

	"github.com/mattn/go-pipeline"
	"github.com/prometheus/common/log"
	"github.com/sawadashota/gh-issue/eloquent"
	"github.com/sawadashota/gh-issue/envchain"
	"github.com/spf13/cobra"
)

const (
	EnvchainNamespace = "gh-issue"
	EnvchainEnv       = "GITHUB_TOKEN"
)

var Set = &cobra.Command{
	Use:   "set",
	Short: "Store GitHub token to envchain",
	Long: `Store GitHub token to envchain
  https://github.com/sorah/envchain

  You can check current env following command.
    $ envchain gh-issue env | grep GITHUB_TOKEN`,
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if err := envchain.Executable(); err != nil {
			eloquent.NewError(err.Error()).Exec()
			os.Exit(1)
		}

		if token == "" {
			eloquent.NewError("Token should be present\n").Exec()
			log.Fatal(cmd.Help())
		}

		execSet(token)
		eloquent.NewSuccess("Set GitHub token successfully\n\n").Exec()
		eloquent.NewSuccess("You can check env following.\n").Exec()
		eloquent.NewSuccess("  $ envchain gh-issue env | grep GITHUB_TOKEN.\n").Exec()
	},
}

func execSet(gitHubToken string) {
	_, err := pipeline.Output(
		[]string{"echo", gitHubToken},
		[]string{"envchain", "--set", "--no-require-passphrase", EnvchainNamespace, EnvchainEnv},
	)

	if err != nil {
		log.Fatal(err)
	}
}
