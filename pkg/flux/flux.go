package flux

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/fluxcd/pkg/recorder"
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

// Handler for Hello request. Should be use with net/http
func (a App) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var event recorder.Event
		if err := httpjson.Parse(r, &event); err != nil {
			httperror.InternalServerError(w, fmt.Errorf("unable to parse event: %s", err))
			return
		}

		w.WriteHeader(http.StatusOK)
		a.send(context.Background(), strings.TrimSpace(event.Message), event.InvolvedObject.Kind, event.InvolvedObject.Namespace, event.InvolvedObject.Name, event.Severity)
	})
}

func (a App) send(ctx context.Context, text string, tags ...string) {
	if strings.HasPrefix(text, "no update") || len(text) > 255 {
		return
	}

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
