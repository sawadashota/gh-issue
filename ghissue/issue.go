package ghissue

type Issues struct {
	Token  string
	Issues []Issue
}

type Issue struct {
	Title    string  `json:"title"`
	Assignee string  `json:"assignee"`
	Labels   []Label `json:"labels"`
}

type Label struct {
	Name string
}

type Option func(issue *Issue)

func New(token string) *Issues {
	return &Issues{Token: token}
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

// Add Labels Option
func WithLabels(labels []Label) Option {
	return func(issue *Issue) {
		issue.Labels = labels
	}
}

// Add Label Option
func WithLabel(label string) Option {
	return func(issue *Issue) {
		issue.Labels = append(issue.Labels, Label{label})
	}
}
