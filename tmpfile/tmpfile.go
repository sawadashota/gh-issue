package tmpfile

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type TmpFile struct {
	editor   string
	filepath string
}

type CreateFileHandler func() error

var (
	stdin  io.Reader
	stdout io.Writer
)

func init() {
	stdin = os.Stdin
	stdout = os.Stdout
}

func New(editor, path string) *TmpFile {
	return &TmpFile{
		editor:   editor,
		filepath: path,
	}
}

func (tf *TmpFile) Open(handler CreateFileHandler) error {
	defer os.Remove(tf.filepath)

	cmd := exec.Command(tf.editor, tf.filepath)

	cmd.Stdin = stdin
	cmd.Stdout = stdout

	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd.Wait()

	if _, err := os.Stat(tf.filepath); err != nil {
		return fmt.Errorf("canceled create file: %s", tf.filepath)
	}

	if err = handler(); err != nil {
		return err
	}

	return nil
}
