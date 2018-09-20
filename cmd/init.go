package cmd

import (
	"os/exec"

	"github.com/sawadashota/gh-issue/issueyaml"

	"github.com/sawadashota/gh-issue/eloquent"
	"github.com/spf13/cobra"
)

const (
	FileName          = "issues.yml"
	IssueFileTemplate = `# template
meta:
  repo: owner/reponame

issues:
  - title: issue title 1
    assignee: assignee
    body: ""
    labels:
      - enhancement`
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create issue file.",
	Long:  `Create issue file.`,
	Run: func(cmd *cobra.Command, args []string) {
		issueyaml.Create(issueyaml.FileName, issueyaml.DefaultTemplate)
		stdoutDependencies()
	},
}

// Stdout guide to install dependencies
func stdoutDependencies() {
	if !executable("envchain") {
		eloquent.NewWarning("Please install envchain", "https://github.com/sorah/envchain").Exec()
	}
}

// Exist command or not
func executable(command string, args ...string) bool {
	err := exec.Command(command, args...).Start()
	return err == nil
}
