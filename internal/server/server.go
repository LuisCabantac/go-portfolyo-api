package server

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/LuisCabantac/go-portfolyo-api/internal/health"
	"github.com/LuisCabantac/go-portfolyo-api/internal/portfolios"
	"github.com/clerk/clerk-sdk-go/v2"
	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	convex "github.com/inselfcontroll/convex-go"
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

func New() *application {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	client := convex.NewClient(os.Getenv("CONVEX_URL"), nil)
	client.SetAdminAuth(os.Getenv("CONVEX_DEPLOY_KEY"))

	clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	return &application{
		config: config{
			addr: ":" + port,
			env:  env,
		},
		logger:     logger,
		client:     client,
		httpClient: &http.Client{},
	}
}

func (app *application) Mount() http.Handler {
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
	r.With(clerkhttp.WithHeaderAuthorization()).Post("/v1/api/portfolios/{portfolioID}/screenshot", portfolioHandler.CreatePortfolioScreenshot)

	return r
}

func (app *application) Run() {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      app.Mount(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	app.logger.Info("starting server", "addr", srv.Addr, "env", app.config.env)

	if err := srv.ListenAndServe(); err != nil {
		app.logger.Error("server failed to start", "error", err)
	}
}
