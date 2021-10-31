package grafana

import (
	"context"
	"flag"
	"strings"

	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
)

type annotationPayload struct {
	Text string
	Tags []string
}

// App of package
type App struct {
	req request.Request
}

// Config of package
type Config struct {
	address  *string
	username *string
	password *string
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string) Config {
	return Config{
		address:  flags.New(prefix, "grafana", "Address").Default("http://grafana", nil).Label("Grafana Address").ToString(fs),
		username: flags.New(prefix, "grafana", "Username").Default("", nil).Label("Grafana Basic Auth Username").ToString(fs),
		password: flags.New(prefix, "grafana", "Password").Default("", nil).Label("Grafana Basic Auth Password").ToString(fs),
	}
}

// New creates new App from Config
func New(config Config) App {
	return App{
		req: request.New().Post(strings.TrimSpace(*config.address)).Path("/api/annotations").BasicAuth(strings.TrimSpace(*config.username), *config.password),
	}
}

// Send grafana annotation
func (a App) Send(ctx context.Context, text string, tags ...string) {
	resp, err := a.req.JSON(ctx, annotationPayload{
		Text: text,
		Tags: tags,
	})
	if err != nil {
		logger.Error("%s", err)
		return
	}

	if err := request.DiscardBody(resp.Body); err != nil {
		logger.Error("unable to discard body: %s", err)
	}
}
