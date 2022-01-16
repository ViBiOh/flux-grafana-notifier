package alertmanager

import (
	"context"
	"flag"
	"net/http"
	"strings"
	"time"

	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	mailer "github.com/ViBiOh/mailer/pkg/client"
	model "github.com/ViBiOh/mailer/pkg/model"
)

type alert struct {
	CommonLabels struct {
		Alertname string `json:"alertname"`
		Service   string `json:"service"`
		Severity  string `json:"severity"`
	} `json:"commonLabels"`
	Receiver    string `json:"receiver"`
	Status      string `json:"status"`
	GroupKey    string `json:"groupKey"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	CommonAnnotations struct {
		Summary string `json:"summary"`
	} `json:"commonAnnotations"`
	ExternalURL string `json:"externalURL"`
	Version     string `json:"version"`
	Alerts      []struct {
		Status string `json:"status"`
		Labels struct {
			Alertname string `json:"alertname"`
			Service   string `json:"service"`
			Severity  string `json:"severity"`
		} `json:"labels"`
		Annotations struct {
			Summary     string `json:"summary"`
			Description string `json:"description"`
		} `json:"annotations"`
		StartsAt     string    `json:"startsAt"`
		EndsAt       time.Time `json:"endsAt"`
		GeneratorURL string    `json:"generatorURL"`
		Fingerprint  string    `json:"fingerprint"`
	} `json:"alerts"`
}

// App of package
type App struct {
	sender    string
	recipient string
	mailerApp mailer.App
}

// Config of package
type Config struct {
	sender    *string
	recipient *string
}

// Flags adds flags for configuring package
func Flags(fs *flag.FlagSet, prefix string, overrides ...flags.Override) Config {
	return Config{
		sender:    flags.New(prefix, "alertmanager", "sender").Default("", overrides).Label("Alertmanager sender").ToString(fs),
		recipient: flags.New(prefix, "alertmanager", "recipient").Default("", overrides).Label("Alertmanager recipient").ToString(fs),
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

		var payload alert
		if err := httpjson.Parse(r, &payload); err != nil {
			httperror.BadRequest(w, err)
			return
		}

		subject := payload.CommonLabels.Alertname
		if payload.Status == "resolved" {
			subject = "[RESOLVED] " + subject
		}

		switch r.URL.Path {
		case "/mail":
			w.WriteHeader(http.StatusOK)
			if err := a.mailerApp.Send(context.Background(), model.NewMailRequest().From(a.sender).As("Alertmanager").WithSubject(subject).To(a.recipient).Template("alertmanager").Data(payload)); err != nil {
				logger.Error("unable to send alertmanager mail: %s", err)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
