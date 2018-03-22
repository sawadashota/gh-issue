package cmd

import "github.com/spf13/cobra"

var (
	file  string
	owner string
	repo  string
	token string
)

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{Use: "gh-issue"}

	// flags for Init

	// flags for Create
	Create.Flags().StringVarP(&owner, "owner", "o", "", "Target repository's owner")
	Create.Flags().StringVarP(&repo, "repo", "r", "", "Target repository")
	Create.Flags().StringVarP(&file, "file", "f", "", "Path for issueYaml file.")

	// flags for Set
	Set.Flags().StringVarP(&token, "token", "t", "", "GitHub Token")

	rootCmd.AddCommand(InitCmd, Create, Set)
	return rootCmd
}
