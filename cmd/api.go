package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/LuisCabantac/portfolyo-go-api/internal/health"
	"github.com/LuisCabantac/portfolyo-go-api/internal/portfolios"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/inselfcontroll/convex-go/src/codebase"
)

type application struct {
	config     config
	logger     *slog.Logger
	client     *codebase.Client
	httpClient *http.Client
}

type config struct {
	addr string
	env  string
}

const apiVersion = "v1"

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	healthService := health.NewService(app.config.env, app.logger)
	healthHandler := health.NewHandler(healthService)
	r.Get("/", healthHandler.HealthCheck)

	portfolioService := portfolios.NewService(app.client, app.httpClient)
	portfolioHandler := portfolios.NewHandler(portfolioService)
	r.With(clerkhttp.WithHeaderAuthorization()).Post(fmt.Sprintf("/%s/api/portfolios/{portfolioID}/screenshot", apiVersion), portfolioHandler.CreatePortfolioScreenshot)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	return srv.ListenAndServe()
}
