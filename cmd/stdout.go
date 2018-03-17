package cmd

import (
	"github.com/fatih/color"
)

type Stdout struct {
	color  *color.Color
	format string
	a      []interface{}
}

type Writer interface {
	Url() Writer
	Important() Writer
	Exec()
}

func NewSuccess(format string, a ...interface{}) Writer {
	return &Stdout{
		color:  color.New(color.FgGreen),
		format: format,
		a:      a,
	}
}

func NewError(format string, a ...interface{}) Writer {
	return &Stdout{
		color:  color.New(color.FgRed),
		format: format,
		a:      a,
	}
}

func NewWarning(messages ...string) Writer {
	var format string

	for k, message := range messages {
		if k == 0 {
			format += "\n"
			format += "*******************************************************\n"
			format += "\n"
		}
		format += "  " + message + "\n"
		if k == len(messages)-1 {
			format += "\n"
			format += "*******************************************************\n"
			format += "\n"

		}
	}

	return &Stdout{
		color:  color.New(color.FgYellow),
		format: format,
	}
}

func (s *Stdout) Important() Writer {
	s.color.Add(color.Bold)
	return s
}

func (s *Stdout) Url() Writer {
	s.color.Add(color.Underline)
	return s
}

func (s *Stdout) Exec() {
	s.color.PrintfFunc()(s.format, s.a...)
}
