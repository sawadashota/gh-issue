package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"github.com/prometheus/common/log"
	"fmt"
	"github.com/fatih/color"
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
		if existIssueFile() {
			log.Fatal(FileName + " is already exists.")
		}

		file, err := os.OpenFile(FileName, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintln(file, IssueFileTemplate)
		color.Cyan("Created " + FileName)
	},
}

func existIssueFile() bool {
	_, err := os.Stat(FileName)
	return err == nil
}
