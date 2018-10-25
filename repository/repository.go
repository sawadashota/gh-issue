// repository package deal with about GitHub repository
package repository

import (
	"context"
	"net/http"

	"github.com/google/go-github/github"

	"golang.org/x/oauth2"
)

type GithubClient struct {
	httpclient *http.Client
	ctx        context.Context
	token      string
}

type OptionClient func(*GithubClient)

func NewClient(ctx context.Context, token string, opts ...OptionClient) *GithubClient {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client := oauth2.NewClient(ctx, ts)

	gc := &GithubClient{
		httpclient: client,
		ctx:        ctx,
		token:      token,
	}

	for _, opt := range opts {
		opt(gc)
	}

	return gc
}

func OptionHTTPClient(c *http.Client) OptionClient {
	return func(gc *GithubClient) {
		gc.httpclient = c
	}
}

func (gc *GithubClient) List() ([]string, error) {
	page := 1
	c := github.NewClient(gc.httpclient)

	var grs []*github.Repository
	for {
		grsop, resp, err := gc.listOfPage(c, page)
		if err != nil {
			return nil, err
		}

		grs = append(grs, grsop...)

		if resp.LastPage == resp.NextPage {
			break
		}

		page++
	}

	var repos []string
	for _, repo := range grs {
		repos = append(repos, *repo.FullName)
	}

	return repos, nil
}

func (gc *GithubClient) listOfPage(c *github.Client, page int) ([]*github.Repository, *github.Response, error) {
	opts := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{
			Page: page,
		},
	}

	grs, resp, err := c.Repositories.List(gc.ctx, "", opts)
	if err != nil {
		return nil, nil, err
	}

	return grs, resp, nil
}
