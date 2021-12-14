package fibr

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ViBiOh/fibr/pkg/provider"
	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	"github.com/ViBiOh/notifier/pkg/discord"
)

// App of package
type App struct {
	secret []byte
	url    string
}

// Config of package
type Config struct {
	secret *string
	url    *string
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) Config {
	return Config{
		secret: flags.New(prefix, "fibr", "Secret").Default("", overrides).Label("Webhook Secret").ToString(fs),
		url:    flags.New(prefix, "fibr", "URL").Default("https://fibr.vibioh.fr", overrides).Label("Fibr URL").ToString(fs),
	}
}

// New creates new App from Config
func New(config Config) App {
	return App{
		secret: []byte(*config.secret),
		url:    strings.TrimSpace(*config.url),
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

		webhookURL := r.URL.Query().Get("discord")
		if len(webhookURL) == 0 {
			return
		}

		content, err := a.getContent(r)
		if err != nil {
			httperror.BadRequest(w, err)
			return
		}

		w.WriteHeader(http.StatusNoContent)

		if len(content) == 0 {
			return
		}

		if err := discord.Send(context.Background(), webhookURL, content); err != nil {
			logger.Error("unable to send discord: %s", err)
		}
	})
}

func (a App) getContent(r *http.Request) (string, error) {
	var e provider.Event
	if err := httpjson.Parse(r, &e); err != nil {
		return "", fmt.Errorf("unable to parse payload: %s", err)
	}

	var content string
	switch e.Type {
	case provider.AccessEvent:
		content = handleAccess(e)
	case provider.UploadEvent:
		content = a.handleFileEvent(e, "uploaded to")
	case provider.DeleteEvent:
		content = a.handleFileEvent(e, "deleted from")
	}

	return content, nil
}

func handleAccess(e provider.Event) string {
	content := strings.Builder{}
	content.WriteString(fmt.Sprintf("\n💻 Someone connected to fibr at %s", e.Time.Format(time.RFC3339)))

	if len(e.Metadata) > 0 {
		content.WriteString("```\n")

		for key, value := range e.Metadata {
			content.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}

		content.WriteString("```")
	}

	return content.String()
}

func (a App) handleFileEvent(e provider.Event, name string) string {
	return fmt.Sprintf("\n💾 Someone %s fibr at %s: %s%s", name, e.Time.Format(time.RFC3339), a.url, e.URL)
}
