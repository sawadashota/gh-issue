package ghissue

import "net/http"

type GitHub struct {
	httpclient *http.Client
	owner      string
	repo       string
	token      string
	issues     []Issue
}

type Option func(*GitHub)

func New(owner string, repo string, token string, opts ...Option) *GitHub {
	hub := &GitHub{
		httpclient: http.DefaultClient,
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

// deprecated: use GitHub
type Issues struct {
	// deprecated: use GitHub
	Owner string
	// deprecated: use GitHub
	Repo string
	// deprecated: use GitHub
	Token  string
	Issues []Issue
}

// deprecated: use New
func NewIssues(owner string, repo string, token string) *Issues {
	return &Issues{
		Owner: owner,
		Repo:  repo,
		Token: token,
	}
}

type Issue struct {
	Title    string
	Assignee string
	Body     string
	Labels   []string
}

type IssueOption func(*Issue)

func (i *Issues) AddIssue(title string, opts ...IssueOption) {
	issue := &Issue{
		Title: title,
	}

	for _, opt := range opts {
		opt(issue)
	}

	i.Issues = append(i.Issues, *issue)
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
