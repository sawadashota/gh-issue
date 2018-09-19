package cmd

import (
	"fmt"
	"os"
	"os/exec"

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
		createIssueFile(FileName, IssueFileTemplate)
		stdoutDependencies()
	},
}

// Create template yaml
func createIssueFile(path, template string) error {
	if existIssueFile(path) {
		return fmt.Errorf("%v is already exists.\n", path)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		return err
	}

	fmt.Fprintln(f, template)
	eloquent.NewSuccess("Created %v successfully\n", path).Exec()

	return nil
}

func existIssueFile(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
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
