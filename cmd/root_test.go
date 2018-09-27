package cmd

import (
	"testing"
)

func TestContains(t *testing.T) {
	cases := map[string]struct {
		data         interface{}
		key          string
		expect       bool
		errorMessage string
	}{
		"contain": {
			data:         map[string]interface{}{"key": "value"},
			key:          "key",
			expect:       true,
			errorMessage: "data doesn't have \"key\" but expected have",
		},
		"not contain": {
			data:         map[string]interface{}{"key": "value"},
			key:          "bar",
			expect:       false,
			errorMessage: "data has \"bar\" but expected not to have",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if contains(c.data, c.key) != c.expect {
				t.Error(c.errorMessage)
			}
		})
	}
}
