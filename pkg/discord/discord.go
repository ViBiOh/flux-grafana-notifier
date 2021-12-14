package discord

import (
	"context"
	"fmt"

	"github.com/ViBiOh/httputils/v4/pkg/request"
)

type discordPayload struct {
	Content string `json:"content"`
}

// Send message to discord webhook
func Send(ctx context.Context, url, content string) error {
	resp, err := request.New().Post(url).JSON(ctx, discordPayload{
		Content: content,
	})
	if err != nil {
		return fmt.Errorf("unable to send discord webhook: %s", err)
	}

	if err = request.DiscardBody(resp.Body); err != nil {
		return fmt.Errorf("unable to discard body: %s", err)
	}

	return nil
}
