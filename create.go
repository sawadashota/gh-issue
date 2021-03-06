package ghissue

import (
	"context"
	"time"

	"github.com/google/go-github/github"
	"github.com/sawadashota/gh-issue/eloquent"
)

type issueCreator func(issueRequest *github.IssueRequest) (*github.Issue, *github.Response, error)

func (gh *GitHub) IssueCreate() *[]Result {
	var results []Result

	issueCreator := client(github.NewClient(gh.httpclient), gh.ctx, gh.owner, gh.repo)

	for k, issue := range gh.issues {
		// GitHub allows to call api 5000 times per hour
		if k != 0 {
			time.Sleep(1 * time.Second)
		}

		results = append(results, *issue.Create(issueCreator))
	}

	return &results
}

// Create client for creating GitHub issue
func client(client *github.Client, ctx context.Context, owner string, repo string) issueCreator {
	return func(issueRequest *github.IssueRequest) (*github.Issue, *github.Response, error) {
		return client.Issues.Create(ctx, owner, repo, issueRequest)
	}
}

// Call GitHub API and create issue
func (i *Issue) Create(issueCreator issueCreator) *Result {
	result := &Result{Issue: *i}

	issueRequest := &github.IssueRequest{
		Title: &i.Title,
	}

	if len(i.Labels) > 0 {
		issueRequest.Labels = &i.Labels
	}

	if i.Assignee != "" {
		issueRequest.Assignee = &i.Assignee
	}

	if i.Body != "" {
		issueRequest.Body = &i.Body
	}

	issueInfo, _, err := issueCreator(issueRequest)

	if err != nil {
		result.Error = err
		eloquent.NewError("\nFailed to create issue.\n%v\n", err.Error()).Exec()
		return result
	}

	result.ApiURL = issueInfo.URL
	eloquent.NewSuccess("\nCreate issue Successfully!\n").Exec()
	eloquent.NewSuccess("%v\n", result.BrowserURL()).Url().Exec()

	return result
}
