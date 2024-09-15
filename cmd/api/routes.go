package api

import (
	"html/template"
	"net/http"
)

func renderReact(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("./web/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	parsedTemplate.Execute(w, nil)
}

func (app *app) Routes() http.Handler {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", fs))

	mux.HandleFunc("GET /{$}", renderReact)
	mux.HandleFunc("GET /about", renderReact)

	return mux
}
