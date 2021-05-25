package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/justinas/alice"
	"github.com/okteto/catalog/api"
	"github.com/okteto/divert"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

var log = zerolog.New(os.Stdout).With().
	Timestamp().
	Logger()

type appCfg struct {
	HTTPPort           string `env:"HTTP_PORT" envDefault:"8080"`
	HealthCheckerURL   string `env:"HEALTH_CHECKER_URL,required"`
	OwnerRegistryURL   string `env:"OWNER_REGISTRY_URL,required"`
	ServiceRegistryURL string `env:"SERVICE_REGISTRY_URL,required"`
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
		divert.InjectDivertHeader(),
	)

	httpClient := http.Client{}
	apiHandler := api.APIHandler{
		HealthClient: api.HealthClient{
			URL:        app.HealthCheckerURL,
			HTTPClient: httpClient,
		},
		OwnerRegistrationClient: api.OwnerRegistrationClient{
			URL:        app.OwnerRegistryURL,
			HTTPClient: httpClient,
		},
		ServiceRegistrationClient: api.ServiceRegistrationClient{
			URL:        app.ServiceRegistryURL,
			HTTPClient: httpClient,
		},
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "health check passed")
	})
	handler.Handle("/data", chain.ThenFunc(apiHandler.Handle))

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
