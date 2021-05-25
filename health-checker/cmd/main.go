package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/justinas/alice"
	"github.com/okteto/catalog/health"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

var log = zerolog.New(os.Stdout).With().
	Timestamp().
	Logger()

type appCfg struct {
	HTTPPort              string `env:"HTTP_PORT" envDefault:"8080"`
	AdvancedHealthEnabled bool   `env:"ADVANCED_HEALTH_ENABLED" envDefault:"false"`
}

func main() {
	app := appCfg{}
	err := env.Parse(&app)
	if err != nil {
		log.Fatal().Err(err).Msg("service is misconfigured")
	}

	chain := alice.New(
		hlog.NewHandler(log),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("m", r.Method).
				Str("u", r.URL.String()).
				Int("s", status).
				Int("z", size).
				Dur("d", duration).
				Msg("")
		}),
		hlog.RemoteAddrHandler("ip"),
		hlog.UserAgentHandler("ua"),
		hlog.RefererHandler("ref"),
		hlog.RequestIDHandler("rid", "Request-Id"),
	)

	var healthClient health.HealthClient
	if app.AdvancedHealthEnabled {
		healthClient = &health.AdvancedHealthClient{}
	} else {
		healthClient = &health.SimpleHealthClient{}
	}

	healthHandler := health.Handler{
		HealthClient: healthClient,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "health check passed")
	})
	handler.Handle("/service-health", chain.ThenFunc(healthHandler.Handle))

	// Start the server and set it up for graceful shutdown.
	srv := &http.Server{
		// It's important to set http server timeouts for the publicly available service api.
		// 60 seconds between when connection is accepted to when the body is fully reaad.
		ReadTimeout: 60 * time.Second,
		// 60 seconds from end of request headers read to end of response write.
		WriteTimeout: 60 * time.Second,
		// 120 seconds for an idle KeeP-Alive connection.
		IdleTimeout: 120 * time.Second,
		Addr:        fmt.Sprintf(":%s", app.HTTPPort),
		Handler:     handler,
	}

	log.Info().Msgf("listening on :%s", app.HTTPPort)
	srv.ListenAndServe()
	log.Info().Msg("shutting server down")
}
