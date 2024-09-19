package api

import (
	"net/http"

	"github.com/dorianneto/burn-secret/internal/handlers"
	"github.com/dorianneto/burn-secret/internal/middleware"
)

func (app *app) Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", fs))

	frontendRoutes := []string{
		"GET /{$}",
		"GET /secret/new",
		"GET /secret/{id}/reveal",
	}

	for _, route := range frontendRoutes {
		mux.HandleFunc(route, handlers.RenderReact)
	}

	secretHandlers := handlers.NewSecretHandlers(app.database)
	userHandlers := handlers.NewUserHandlers(app.database, app.logger)
	authHandlers := handlers.NewAuthHandlers(app.database, app.logger)

	mux.HandleFunc("GET /api/v1/secret/{id}", secretHandlers.ShowSecret)
	mux.HandleFunc("POST /api/v1/secret/new", secretHandlers.GenerateSecret)
	mux.HandleFunc("DELETE /api/v1/secret/{id}/burn", secretHandlers.BurnSecret)

	mux.HandleFunc("POST /api/v1/user/new", userHandlers.RegisterUser)

	mux.HandleFunc("POST /api/v1/login", authHandlers.Login)

	return middleware.LogRequests(mux, app.logger)
}
