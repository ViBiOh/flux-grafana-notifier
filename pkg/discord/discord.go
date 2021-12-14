package discord

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

type discordPayload struct {
	Content string `json:"content"`
}

// App of package
type App struct {
	webhookReq  request.Request
	cyclismeReq request.Request
}

// Config of package
type Config struct {
	webhookURL  *string
	cyclismeURL *string
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) Config {
	return Config{
		webhookURL:  flags.New(prefix, "discord", "WebhookURL").Default("WebhookURL", overrides).Label("Discord WebhookURL").ToString(fs),
		cyclismeURL: flags.New(prefix, "discord", "CyclismeURL").Default("CyclismeURL", overrides).Label("Cyclisme Discord WebhookURL").ToString(fs),
	}
}

// New creates new App from Config
func New(config Config) App {
	return App{
		webhookReq:  request.New().Post(strings.TrimSpace(*config.webhookURL)),
		cyclismeReq: request.New().Post(strings.TrimSpace(*config.cyclismeURL)),
	}
}

// Enabled checks that requirements are met
func (a App) Enabled() bool {
	return !a.webhookReq.IsZero() || !a.cyclismeReq.IsZero()
}

// Send discord webhook content
func (a App) Send(ctx context.Context, content string) error {
	resp, err := a.webhookReq.JSON(ctx, discordPayload{
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

// SendCyclisme messae to cyclisme discord webhook content
func (a App) SendCyclisme(ctx context.Context, content string) error {
	resp, err := a.cyclismeReq.JSON(ctx, discordPayload{
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
