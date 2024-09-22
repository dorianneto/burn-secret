package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/dorianneto/burn-secret/cmd/api"
	"github.com/dorianneto/burn-secret/internal"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env." + os.Getenv("APP_ENV"))

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	database, err := internal.NewDatabase(logger)
	if err != nil {
		logger.Error(err.Error())
	}

	app := api.NewApp(logger, database)

	logger.Info("server running on port :80")

	if err := http.ListenAndServe(":80", app.Routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
