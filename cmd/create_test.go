package cmd

import (
	"testing"
	"regexp"
)

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

func TestGetToken(t *testing.T) {
	token, err := getToken()
	if err != nil {
		t.Error(err)
	}

	r := regexp.MustCompile("^.+$")
	if !r.MatchString(token) {
		t.Errorf("Invalid Token: %v\n", token)
	}
}
