package cmd

import "testing"

func TestExecutable(t *testing.T) {
	if !executable("ls") {
		t.Error("executable(\"ls\") should be true")
	}
}

func TestExecutable02(t *testing.T) {
	if executable("hogehoge") {
		t.Error("executable(\"hogehoge\") should be false")
	}
}
