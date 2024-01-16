package main

import (
	"fmt"
	"os"
	"log/slog"
	"context"
	"net/http"
	"io"
	"net/url"
	"strings"
	"strconv"

	"github.com/google/go-github/v58/github"
	"github.com/wzshiming/gh-gpt/pkg/run"
	"github.com/wzshiming/gh-pr-summarize/source"
	"github.com/wzshiming/gh-pr-summarize/prompts"
)

func usage() {
	fmt.Println("Usage: gh-pr-summarize <pr url>")
	os.Exit(1)
}

func parseURL(uri string) (owner, repo string, number int, err error) {
	i, err := url.Parse(uri)
	if err != nil {
		return "", "", 0, err
	}

	if i.Host != "github.com" {
		return "", "", 0, fmt.Errorf("host %s", i.Host)
	}

	paths := strings.Split(i.Path, "/")
	if len(paths) != 5 {
		return "", "", 0, fmt.Errorf("path %s", i.Path)
	}
	if paths[3] != "pull" {
		return "", "", 0, fmt.Errorf("path %s", i.Path)
	}

	n := paths[4]
	if i := strings.Index(n, "."); i != -1 {
		n = n[:i]
	}
	number, err = strconv.Atoi(n)
	if err != nil {
		return "", "", 0, fmt.Errorf("path %s", i.Path)
	}
	owner = paths[1]
	repo = paths[2]
	return owner, repo, number, nil
}

func main() {
	ctx := context.Background()

	if len(os.Args) <= 1 {
		usage()
	}

	owner, repo, number, err := parseURL(os.Args[1])
	if err != nil {
		slog.Error("parse url", "err", err)
		os.Exit(1)
	}

	client := source.NewClient(os.Getenv("GH_TOKEN"))

	var (
		title      string
		discussion []prompts.Comment
	)

	patch, err := getPatch(owner, repo, number)
	if err != nil {
		slog.Error("get patch", "err", err)
		os.Exit(1)
	}

	issue, err := client.GetIssue(ctx, owner, repo, number)
	if err != nil {
		slog.Error("get issue", "err", err)
		os.Exit(1)
	}

	title = *issue.Title
	discussion = append(discussion, prompts.Comment{
		User: *issue.User.Login,
		Body: defaultValue(issue.Body),
	})

	err = client.ListCommits(ctx, owner, repo, number, func(comment *github.IssueComment) error {
		discussion = append(discussion, prompts.Comment{
			User: *comment.User.Login,
			Body: defaultValue(comment.Body),
		})
		return nil
	})
	if err != nil {
		slog.Error("list commits", "err", err)
		os.Exit(1)
	}

	summary, err := run.Run(ctx, prompts.ConventionalPullRequest(title, discussion, patch))
	if err != nil {
		slog.Error("conventional PR", "err", err)
		os.Exit(1)
	}

	summary = strings.Replace(summary, "\n\n", "\n", -1)
	summary = strings.TrimSpace(summary)

	fmt.Printf("%s\n", summary)
}

func getPatch(owner, repo string, number int) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://github.com/%s/%s/pull/%d.diff", owner, repo, number))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code %d", resp.StatusCode)
	}

	patch, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(patch), nil
}

func defaultValue[T any](t *T) T {
	var z T
	if t == nil {
		return z
	}
	return *t
}
