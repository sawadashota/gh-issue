package cmd

import (
	"github.com/spf13/cobra"
	"github.com/prometheus/common/log"
	"os/exec"
)

const (
	EnvchainNamespace = "gh-issue"
	EnvchainEnv       = "GITHUB_TOKEN"
)

var Set = &cobra.Command{
	Use:   "set",
	Short: "Store GitHub token to envchain",
	Long: `Store GitHub token to envchain
  https://github.com/sorah/envchain`,
	Args: cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if !executable("envchain") {
			NewWarning("Please install envchain", "https://github.com/sorah/envchain").Exec()
			log.Fatal("Command envchain is not exists.")
		}

		if token == "" {
			NewError("Token should be present\n").Exec()
			log.Fatal(cmd.Help())
		}


		command := exec.Command("envchain", "--set", "--no-require-passphrase", EnvchainNamespace, EnvchainEnv)
		_, err := command.Output()
		if err != nil {
			log.Fatal(err)
		}

		println(token)
	},
}
