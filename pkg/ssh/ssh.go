package ssh

import (
	"context"
	"flag"
	"net/http"
	"strings"

	"github.com/ViBiOh/httputils/v4/pkg/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/request"
	mailer "github.com/ViBiOh/mailer/pkg/client"
	model "github.com/ViBiOh/mailer/pkg/model"
)

// App of package
type App struct {
	mailerApp mailer.App
	sender    string
	recipient string
}

// Config of package
type Config struct {
	sender    *string
	recipient *string
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) Config {
	return Config{
		sender:    flags.New(prefix, "ssh", "sender").Default("", overrides).Label("SSH Notification sender").ToString(fs),
		recipient: flags.New(prefix, "ssh", "recipient").Default("", overrides).Label("SSH Notification recipient").ToString(fs),
	}
}

// New creates new App from Config
func New(config Config, mailerApp mailer.App) App {
	return App{
		sender:    strings.TrimSpace(*config.sender),
		recipient: strings.TrimSpace(*config.recipient),
		mailerApp: mailerApp,
	}
}

// Handler for Hello request. Should be use with net/http
func (a App) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		content, err := request.ReadBodyRequest(r)
		if err != nil {
			httperror.InternalServerError(w, err)
			return
		}

		switch r.URL.Path {
		case "/mail":
			w.WriteHeader(http.StatusOK)
			if err := a.mailerApp.Send(context.Background(), model.NewMailRequest().From(a.sender).As("SSH Monitoring").To(a.recipient).Template("ssh").Data(content)); err != nil {
				logger.Error("unable to send ssh mail: %s", err)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
