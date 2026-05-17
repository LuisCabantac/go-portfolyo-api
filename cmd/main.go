package main

import (
	"log/slog"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	cfg := config{
		addr: ":" + port,
		env:  env,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	api := application{
		config: cfg,
		logger: logger,
	}

	if err := api.run(api.mount()); err != nil {
		api.logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
