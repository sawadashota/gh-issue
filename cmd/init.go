package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"github.com/prometheus/common/log"
	"fmt"
	"os/exec"
)

const (
	FileName          = "issues.yml"
	IssueFileTemplate = `issues:
  - title: issue title 1
    assignee: assignee
    labels:
      - enhancement`
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Create issue file.",
	Long:  `Create issue file.`,
	Run: func(cmd *cobra.Command, args []string) {
		createIssueFile()
		stdoutDependencies()
	},
}

// Create template yaml
func createIssueFile() {
	if existIssueFile() {
		NewError("%v is already exists.\n", FileName).Exec()
		os.Exit(1)
	}

	file, err := os.OpenFile(FileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(file, IssueFileTemplate)
	NewSuccess("Created %v successfully\n", FileName).Exec()
}

func existIssueFile() bool {
	_, err := os.Stat(FileName)
	return err == nil
}

// Stdout guide to install dependencies
func stdoutDependencies() {
	if !executable("envchain") {
		NewWarning("Please install envchain", "https://github.com/sorah/envchain").Exec()
	}
}

// Exist command or not
func executable(command string, args ...string) bool {
	err := exec.Command(command, args...).Start()
	return err == nil
}
