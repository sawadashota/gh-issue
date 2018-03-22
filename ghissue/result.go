package ghissue

import "regexp"

type Result struct {
	Issue  *Issue
	ApiURL *string
	Error  error
}

func(r *Result) BrowserURL() string {
	regex := regexp.MustCompile(`^https:\/\/api\.github\.com\/repos\/`)
	return regex.ReplaceAllString(*r.ApiURL, "https://github.com/")
}
