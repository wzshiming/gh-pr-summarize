package source

import (
	"context"

	"github.com/google/go-github/v58/github"
)

func (c *Client) GetIssue(ctx context.Context, owner, repo string, number int) (*github.Issue, error) {
	i, _, err := c.Client.Issues.Get(ctx, owner, repo, number)
	return i, err
}
