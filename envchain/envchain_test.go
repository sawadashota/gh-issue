package envchain_test

import (
	"regexp"
	"testing"

	"github.com/sawadashota/gh-issue/envchain"
)

var tokenRegexp *regexp.Regexp

func init() {
	tokenRegexp = regexp.MustCompile("^.+$")
}

func TestToken(t *testing.T) {
	token, err := envchain.Token()
	if err != nil {
		t.Error(err)
	}

	if !tokenRegexp.MatchString(token) {
		t.Errorf("Invalid Token: %v\n", token)
	}
}
