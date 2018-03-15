package cmd

import "github.com/spf13/cobra"

var (
	file string
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{Use: "gh-issue"}

	// flags for init

	// flags for create
	Create.Flags().StringVarP(&file, "file", "f", "", "Path for issueYaml file.")

	rootCmd.AddCommand(InitCmd, Create)
	return rootCmd
}
