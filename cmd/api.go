package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/LuisCabantac/portfolyo-go-api/internal/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
	logger *slog.Logger
}

type config struct {
	addr string
	env  string
}

const apiName = "Portfolyo API"

const version = "1.0.0"

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if err := json.Write(w, http.StatusOK, struct {
			Message     string
			Version     string
			Status      string
			Environment string
		}{
			Message:     apiName,
			Version:     version,
			Status:      "running",
			Environment: app.config.env,
		}); err != nil {
			app.logger.Error("the server encountered a problem and could not process your request", "error", err)
		}
	})

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
