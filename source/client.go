package source

import (
	"net/http"
	"context"

	"github.com/google/go-github/v58/github"
	"golang.org/x/oauth2"
)

var maxPerPage = 100

type Client struct {
	Client *github.Client
}

func NewClient(token string) *Client {
	var c *http.Client

	if token != "" {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		c = oauth2.NewClient(ctx, ts)
	}

	return &Client{
		Client: github.NewClient(c),
	}
}
