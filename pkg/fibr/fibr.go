package fibr

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/notifier/pkg/discord"
)

type event struct {
	Metadata map[string]string `json:"metadata"`
	Type     string            `json:"type"`
}

// App of package
type App struct {
	discordApp discord.App
	secret     []byte
}

// Config of package
type Config struct {
	secret *string
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) Config {
	return Config{
		secret: flags.New(prefix, "fibr", "Secret").Default("", overrides).Label("Webhook Secret").ToString(fs),
	}
}

// New creates new App from Config
func New(config Config, discordApp discord.App) App {
	return App{
		discordApp: discordApp,
		secret:     []byte(*config.secret),
	}
}

// Handler for Hello request. Should be use with net/http
func (a App) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if len(a.secret) != 0 {
			if ok, err := request.ValidateSignature(r, a.secret); err != nil {
				httperror.BadRequest(w, err)
				return
			} else if !ok {
				httperror.Unauthorized(w, errors.New("signature invalid"))
				return
			}
		}

		if !a.discordApp.Enabled() {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var e event
		if err := httpjson.Parse(r, &e); err != nil {
			httperror.BadRequest(w, err)
			return
		}

		if e.Type != "access" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		content := strings.Builder{}
		content.WriteString(fmt.Sprintf("Someone connected to fibr at %s", time.Now().Format(time.RFC3339)))

		if len(e.Metadata) > 0 {
			content.WriteString("```\n")

			for key, value := range e.Metadata {
				content.WriteString(fmt.Sprintf("%s: %s\n", key, value))
			}

			content.WriteString("```")
		}

		switch r.URL.Path {
		case "/fibr/discord":
			w.WriteHeader(http.StatusNoContent)
			if err := a.discordApp.Send(context.Background(), content.String()); err != nil {
				logger.Error("unable to send discord: %s", err)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
