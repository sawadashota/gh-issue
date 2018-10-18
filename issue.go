package ghissue

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type GitHub struct {
	httpclient *http.Client
	ctx        context.Context
	owner      string
	repo       string
	token      string
	issues     []Issue
}

type Option func(*GitHub)

func New(ctx context.Context, owner string, repo string, token string, opts ...Option) *GitHub {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client := oauth2.NewClient(ctx, ts)

	hub := &GitHub{
		httpclient: client,
		ctx:        ctx,
		owner:      owner,
		repo:       repo,
		token:      token,
	}

	for _, opt := range opts {
		opt(hub)
	}

	return hub
}

func OptionHTTPClient(c *http.Client) Option {
	return func(hub *GitHub) {
		hub.httpclient = c
	}
}

type Issue struct {
	Title    string
	Assignee string
	Body     string
	Labels   []string
}

type IssueOption func(*Issue)

func (gh *GitHub) AddIssue(title string, opts ...IssueOption) {
	issue := &Issue{
		Title: title,
	}

	for _, opt := range opts {
		opt(issue)
	}

	gh.issues = append(gh.issues, *issue)
}

// Add Assignee IssueOption
func WithAssignee(assignee string) IssueOption {
	return func(issue *Issue) {
		issue.Assignee = assignee
	}
}

// Add Body IssueOption
func WithBody(body string) IssueOption {
	return func(issue *Issue) {
		issue.Body = body
	}
}

// Add Label IssueOption
func WithLabel(label string) IssueOption {
	return func(issue *Issue) {
		issue.Labels = append(issue.Labels, label)
	}
}
