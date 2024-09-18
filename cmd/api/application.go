package api

import (
	"log/slog"

	"github.com/dorianneto/burn-secret/internal/interfaces"
)

type app struct {
	logger   *slog.Logger
	database interfaces.KeyPairBased
}

func NewApp(logger *slog.Logger, database interfaces.KeyPairBased) *app {
	return &app{
		logger:   logger,
		database: database,
	}
}
