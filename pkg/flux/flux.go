package flux

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/httpjson"
	"github.com/ViBiOh/notifier/pkg/grafana"
	"github.com/fluxcd/pkg/recorder"
)

// App of package
type App struct {
	grafanaApp grafana.App
}

// New creates new App from Config
func New(grafanaApp grafana.App) App {
	return App{
		grafanaApp: grafanaApp,
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

		text := strings.TrimSpace(event.Message)
		if strings.HasPrefix(text, "no update") || len(text) > 255 {
			return
		}

		switch r.URL.Path {
		case "/grafana":
			w.WriteHeader(http.StatusOK)
			a.grafanaApp.Send(context.Background(), text, event.InvolvedObject.Kind, event.InvolvedObject.Namespace, event.InvolvedObject.Name, event.Severity)
		default:
			w.WriteHeader(http.StatusBadRequest)
		}
	})
}
