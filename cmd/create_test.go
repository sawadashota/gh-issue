package cmd

import "testing"

func TestContains(t *testing.T) {
	var data interface{}
	data = map[string]interface{}{"key": "value"}

	if !contains(data, "key") {
		t.Error("Data doesn't have \"key\"")
	}
}

func TestContains02(t *testing.T) {
	var data interface{}
	data = map[string]interface{}{"key": "value"}

	if contains(data, "bar") {
		t.Error("Data has \"bar\"")
	}
}