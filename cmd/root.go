package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"github.com/sawadashota/gh-issue"
	"github.com/sawadashota/gh-issue/config"
	"github.com/sawadashota/gh-issue/eloquent"
	"github.com/sawadashota/gh-issue/envchain"
	"github.com/sawadashota/gh-issue/issueyaml"
	"github.com/sawadashota/gh-issue/tmpfile"
	"github.com/spf13/cobra"
)

const (
	ConfigDir = "~/.config/gh-issue/"
)

var (
	token string
)

type tomlConfig struct {
	Editor   string `toml:"editor"`
	Template string `toml:"template"`
}

func RootCmd() *cobra.Command {
	// flags for root

	// flags for SetCmd
	SetCmd.Flags().StringVarP(&token, "token", "t", "", "GitHub Token")

	rootCmd.AddCommand(SetCmd, EditCmd, RepoCmd)
	return rootCmd
}

var rootCmd = &cobra.Command{
	Use: "gh-issue",
	Run: func(cmd *cobra.Command, args []string) {
		if err := envchain.Executable(); err != nil {
			eloquent.NewError(err.Error()).Exec()
			os.Exit(1)
		}

		baseDir, err := baseDirAbs(ConfigDir)
		if err != nil {
			log.Fatalln(err)
		}

		issueFilePath := filepath.Join(baseDir, issueyaml.FileName)
		configFilePath := config.Path(baseDir)

		err = config.Generate(configFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		tc, err := readConfig(configFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		// TODO: select issue by peco
		if err := tc.replaceVariable(); err != nil {
			log.Fatalln(err)
		}
		_ = issueyaml.Create(issueFilePath, tc.Template)

		err = tmpfile.New(tc.Editor, issueFilePath).Open(func() error {
			token, err := envchain.Token()
			if err != nil {
				return err
			}

			ctx := context.Background()
			return createIssues(ctx, issueFilePath, token)
		})

		if err != nil {
			log.Fatalln(err)
		}
	},
}

func baseDirAbs(baseDir string) (string, error) {
	baseDir, err := homedir.Expand(baseDir)
	if err != nil {
		return "", err
	}

	if err = os.MkdirAll(baseDir, 0777); err != nil {
		return "", err
	}

	return baseDir, nil
}

func readConfig(path string) (*tomlConfig, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	var conf tomlConfig
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

// TODO: read from ./.git/config
func (tc *tomlConfig) replaceVariable() error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	dir, base := filepath.Split(pwd)

	project := fmt.Sprintf("%s/%s", filepath.Base(dir), base)

	tc.Template = strings.Replace(tc.Template, "${CURRENT_PROJECT}", project, -1)

	return nil
}

func createIssues(ctx context.Context, fp, token string) error {
	y, err := issueyaml.New(fp)

	if err != nil {
		return err
	}

	issues, err := y.Issues(ctx, token)
	if err != nil {
		return err
	}

	results := issues.IssueCreate()

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

	if errCount > 1 {
		eloquent.NewError("%d errors occurred.\n", errCount).Important().Exec()
	}
}

func contains(s interface{}, e string) bool {
	for key := range s.(map[string]interface{}) {
		if e == key {
			return true
		}
	}
	return false
}
