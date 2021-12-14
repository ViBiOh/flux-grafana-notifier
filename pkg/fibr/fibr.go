package fibr

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"path"
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

		content, err := getContent(r)
		if err != nil {
			httperror.BadRequest(w, err)
			return
		}

		if len(content) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		switch r.URL.Path {
		case "/fibr/discord":
			w.WriteHeader(http.StatusNoContent)
			if err := a.discordApp.Send(context.Background(), content); err != nil {
				logger.Error("unable to send discord: %s", err)
			}
		case "/fibr/cyclisme":
			w.WriteHeader(http.StatusNoContent)
			if err := a.discordApp.SendCyclisme(context.Background(), content); err != nil {
				logger.Error("unable to send discord: %s", err)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}

func getContent(r *http.Request) (string, error) {
	var e provider.Event
	if err := httpjson.Parse(r, &e); err != nil {
		return "", fmt.Errorf("unable to parse payload: %s", err)
	}

	var content string
	switch e.Type {
	case provider.AccessEvent:
		content = handleAccess(e)
	case provider.UploadEvent:
		content = handleFileEvent(e, "uploaded to")
	case provider.DeleteEvent:
		content = handleFileEvent(e, "deleted from")
	}

	return content, nil
}

func handleAccess(e provider.Event) string {
	content := strings.Builder{}
	content.WriteString(fmt.Sprintf("\nSomeone connected to fibr at %s", time.Now().Format(time.RFC3339)))

	if len(e.Metadata) > 0 {
		content.WriteString("```\n")

		for key, value := range e.Metadata {
			content.WriteString(fmt.Sprintf("%s: %s\n", key, value))
		}

		content.WriteString("```")
	}

	return content.String()
}

func handleFileEvent(e provider.Event, name string) string {
	content := strings.Builder{}

	content.WriteString(fmt.Sprintf("\n💾 Someone %s fibr at %s", name, time.Now().Format(time.RFC3339)))
	content.WriteString(fmt.Sprintf("\n🗂 Folder `%s`", path.Dir(e.Item.Pathname)))
	content.WriteString(fmt.Sprintf("\n📸 Name `%s`", e.Item.Name))

	return content.String()
}
