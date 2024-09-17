package handlers

import (
	"net/http"
	"text/template"
)

func RenderReact(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("./web/index.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	parsedTemplate.Execute(w, nil)
}
