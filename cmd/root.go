package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sawadashota/gh-issue/issueyaml"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"github.com/sawadashota/gh-issue/tmpfile"
	"github.com/spf13/cobra"
)

const (
	ConfigDir         = "~/.config/gh-issue/"
	TmpIssueFilename  = "issue.yml"
	ConfigFilename    = "config.toml"
	DefaultConfFormat = `editor = "vim"
template = """%s"""`
)

var (
	file  string
	token string
)

type tomlConfig struct {
	Editor   string `toml:"editor"`
	Template string `toml:"template"`
}

func RootCmd() *cobra.Command {
	// flags for root

	// flags for Init

	// flags for Create
	Create.Flags().StringVarP(&file, "file", "f", "", "Path for issueYaml file.")

	// flags for Set
	Set.Flags().StringVarP(&token, "token", "t", "", "GitHub Token")

	rootCmd.AddCommand(InitCmd, Create, Set)
	return rootCmd
}

var rootCmd = &cobra.Command{
	Use: "gh-issue",
	Run: func(cmd *cobra.Command, args []string) {

		baseDir, err := baseDirAbs(ConfigDir)
		if err != nil {
			log.Fatalln(err)
		}

		issueFilePath := filepath.Join(baseDir, TmpIssueFilename)
		configFilePath := filepath.Join(baseDir, ConfigFilename)

		if err != nil {
			log.Fatalln(err)
		}

		err = generateDefaultConfIfNoExists(configFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		tc, err := readConfig(configFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		_ = issueyaml.Create(issueFilePath, tc.Template)

		err = tmpfile.New(tc.Editor, issueFilePath).Open(func() error {
			token, err := getToken()
			if err != nil {
				return err
			}

			return createIssues(issueFilePath, token)
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

func generateDefaultConfIfNoExists(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		return err
	}

	_, err = fmt.Fprintln(f, defaultConf())

	if err != nil {
		return err
	}

	return nil
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

func defaultConf() string {
	return fmt.Sprintf(DefaultConfFormat, IssueFileTemplate)
}
