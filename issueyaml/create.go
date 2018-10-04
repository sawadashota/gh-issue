package issueyaml

import (
	"fmt"
	"os"
)

const (
	FileName = "issues.yml"
)

// Create template yaml
func Create(path, template string) error {
	if exists(path) {
		return fmt.Errorf("%v is already exists.\n", path)
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	defer f.Close()

	if err != nil {
		return err
	}

	fmt.Fprintln(f, template)
	return nil
}

func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil
}
