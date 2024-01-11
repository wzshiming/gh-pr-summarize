package source

import (
	"context"

	"github.com/google/go-github/v58/github"
)

func (c *Client) ListCommits(ctx context.Context, owner, repo string, number int, fn func(commit *github.IssueComment) error) error {
	return c.listComments(ctx, owner, repo, number, fn, 0)
}

func (c *Client) listComments(ctx context.Context, owner, repo string, number int, fn func(comment *github.IssueComment) error, page int) error {
	comments, resp, err := c.Client.Issues.ListComments(ctx, owner, repo, number, &github.IssueListCommentsOptions{
		ListOptions: github.ListOptions{
			Page:    page,
			PerPage: maxPerPage,
		},
	})
	if err != nil {
		return err
	}

	for _, comment := range comments {
		if err := fn(comment); err != nil {
			return err
		}
	}

	if resp.NextPage != 0 {
		return c.listComments(ctx, owner, repo, number, fn, resp.NextPage)
	}
	return nil
}
