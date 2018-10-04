package config

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	Filename          = "config.toml"
	DefaultConfFormat = `editor = "vim"
template = """%s"""`
	DefaultTemplate = `# template
meta:
  repo: owner/reponame

issues:
  - title: issue title 1
    assignee: assignee
    body: ""
    labels:
      - enhancement`
)

func Generate(path string) error {
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

func Path(dir string) string {
	return filepath.Join(dir, Filename)
}

func defaultConf() string {
	return fmt.Sprintf(DefaultConfFormat, DefaultTemplate)
}
