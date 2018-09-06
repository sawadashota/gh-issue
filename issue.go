package ghissue

type Issues struct {
	Owner  string
	Repo   string
	Token  string
	Issues []Issue
}

type Issue struct {
	Title    string
	Assignee string
	Body     string
	Labels   []string
}

type Option func(issue *Issue)

func New(owner string, repo string, token string) *Issues {
	return &Issues{
		Owner: owner,
		Repo:  repo,
		Token: token,
	}
}

func (i *Issues) AddIssue(title string, options ...Option) {
	issue := &Issue{
		Title: title,
	}

	for _, option := range options {
		option(issue)
	}

	i.Issues = append(i.Issues, *issue)
}

// Add Assignee Option
func WithAssignee(assignee string) Option {
	return func(issue *Issue) {
		issue.Assignee = assignee
	}
}

// Add Body Option
func WithBody(body string) Option {
	return func(issue *Issue) {
		issue.Body = body
	}
}

// Add Label Option
func WithLabel(label string) Option {
	return func(issue *Issue) {
		issue.Labels = append(issue.Labels, label)
	}
}
