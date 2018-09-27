package envchain

import (
	"fmt"
	"os/exec"
	"regexp"
)

const (
	Command           = "envchain"
	EnvchainNamespace = "gh-issue"
	EnvchainEnv       = "GITHUB_TOKEN"
)

func Token() (string, error) {
	if err := Executable(); err != nil {
		return "", err
	}

	command := exec.Command(Command, EnvchainNamespace, "env")
	res, err := command.Output()

	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(EnvchainEnv + `=(.+)`)
	matches := r.FindStringSubmatch(string(res))

	if len(matches) < 2 {
		return "", fmt.Errorf("cannot Find GITHUB_TOKEN")
	}

	return matches[1], nil
}

// Exist command or not
func Executable() error {
	err := exec.Command(Command).Start()
	if err != nil {
		return fmt.Errorf("please install envchain https://github.com/sorah/envchain")
	}

	return nil
}
