package gpt

import (
	"fmt"
	"context"
	"errors"
	"os"
	"path/filepath"
	"bytes"
	"strings"

	"github.com/wzshiming/gh-gpt/pkg/api"
	"github.com/wzshiming/gh-gpt/pkg/auth"
	"github.com/wzshiming/gh-gpt/pkg/cache"
)

func Generate(ctx context.Context, content string) (string, error) {
	hosts := auth.Hosts()

	oauth, err := hosts.GetToken()
	if err != nil {
		return "", fmt.Errorf("failed to get oauth token: %w", err)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir: %w", err)
	}

	tokenCachePath := filepath.Join(home, ".gh-gpt/token.json")
	tokenCache := cache.NewFileCache(tokenCachePath)
	cli := api.NewClient(
		api.WithTokenCache(tokenCache),
	)

	token, err := cli.TokenWishCache(ctx, oauth)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}

	req := api.ChatRequest{
		Model: "gpt-4",
		Messages: []api.Message{
			{Role: "user", Content: content},
		},
	}

	buf := bytes.NewBuffer(nil)
	fn := func(resp api.ChatResponse) error {
		for _, choice := range resp.Choices {
			if choice.Delta.Content != "" {
				fmt.Fprint(buf, choice.Delta.Content)
			}
			if choice.Message.Content != "" {
				fmt.Fprintln(buf, choice.Message.Content)
			}
		}
		return nil
	}

	err = cli.ChatCompletions(ctx, token, &req, fn)
	if err != nil {
		if !errors.Is(err, context.Canceled) {
			return "", fmt.Errorf("failed to chat: %w", err)
		}
	}

	return strings.TrimSpace(buf.String()), nil
}
