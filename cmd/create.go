package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/sawadashota/gh-issue/issueyaml"

	"github.com/sawadashota/gh-issue"
	"github.com/sawadashota/gh-issue/eloquent"
	"github.com/spf13/cobra"
)

var Create = &cobra.Command{
	Use:   "create -f [filepath]",
	Short: "Create issue at GitHub",
	Long:  `Create issue at GitHub`,
	Args:  cobra.MaximumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		token, err := getToken()
		if err != nil {
			eloquent.NewError(err.Error()).Exec()
		}

		if err := createIssues(file, token); err != nil {
			eloquent.NewError(err.Error()).Exec()
		}
	},
}

func createIssues(fp, token string) error {
	y, err := issueyaml.New(fp)

	if err != nil {
		return err
	}

	issues, err := y.Issues(token)
	if err != nil {
		return err
	}

	results := issues.Create()

	stdoutAllError(results)

	return nil
}

func stdoutAllError(results *[]ghissue.Result) {
	errCount := 0
	for _, result := range *results {
		if result.HasError() {
			errCount++

			if errCount == 1 {
				eloquent.NewError("\n************* Error List *************\n\n").Important().Exec()
			}

			eloquent.NewError("%d. %v\n%v\n\n", errCount, result.Issue.Title, result.Error.Error()).Exec()
		}
	}

	eloquent.NewError("%d errors occurred.\n", errCount).Important().Exec()
}

func contains(s interface{}, e string) bool {
	for key := range s.(map[string]interface{}) {
		if e == key {
			return true
		}
	}
	return false
}

func getToken() (string, error) {
	if !executable("envchain") {
		eloquent.NewWarning("Please install envchain", "https://github.com/sorah/envchain").Exec()
		log.Fatal("Command envchain is not exists.")
	}

	command := exec.Command("envchain", "gh-issue", "env")
	res, err := command.Output()

	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(EnvchainEnv + `=(.+)`)
	matches := r.FindStringSubmatch(string(res))

	if len(matches) < 2 {
		return "", fmt.Errorf("Cannot Find GITHUB_TOKEN\n")
	}

	return matches[1], nil
}
