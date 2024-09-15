package api

import "log/slog"

type app struct {
	logger *slog.Logger
}

func NewApp(logger *slog.Logger) *app {
	return &app{
		logger: logger,
	}
}
