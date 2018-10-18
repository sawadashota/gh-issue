package cmd

import (
	"io"
	"os"
	"os/exec"

	"github.com/prometheus/common/log"
	"github.com/sawadashota/gh-issue/config"
	"github.com/spf13/cobra"
)

var (
	stdin  io.Reader
	stdout io.Writer
)

func init() {
	stdin = os.Stdin
	stdout = os.Stdout
}

var EditCmd = &cobra.Command{
	Use:   "edit",
	Short: "edit config",
	Run: func(c *cobra.Command, args []string) {
		baseDir, err := baseDirAbs(ConfigDir)
		if err != nil {
			log.Fatalln(err)
		}

		configFilePath := config.Path(baseDir)

		err = config.Generate(configFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		tc, err := readConfig(configFilePath)
		if err != nil {
			log.Fatalln(err)
		}

		if err := edit(tc.Editor, configFilePath); err != nil {
			log.Fatalln(err)
		}
	},
}

func edit(editor, file string) error {
	command := exec.Command(editor, file)

	command.Stdin = stdin
	command.Stdout = stdout

	err := command.Run()
	if err != nil {
		return err
	}
	command.Wait()

	return nil
}
