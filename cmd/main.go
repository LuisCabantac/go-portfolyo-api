package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"
	convex "github.com/inselfcontroll/convex-go"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	cfg := config{
		addr: ":" + port,
		env:  env,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	httpClient := &http.Client{}

	api := application{
		config:     cfg,
		logger:     logger,
		client:     client,
		httpClient: httpClient,
	}

	if err := api.run(api.mount()); err != nil {
		api.logger.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
