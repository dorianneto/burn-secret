package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/dorianneto/burn-secret/cmd/api"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := api.NewApp(logger)

	logger.Info("server running on port :8080")

	if err := http.ListenAndServe(":8080", app.Routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
