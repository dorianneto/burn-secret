package handlers

import (
	"net/http"

	"github.com/dorianneto/burn-secret/internal/utils"
)

type secretHandlers struct{}

func NewSecretHandlers() *secretHandlers {
	return &secretHandlers{}
}

func (sh *secretHandlers) GenerateSecret(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, map[string]string{"page": "new"})
}

func (sh *secretHandlers) BurnSecret(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, map[string]string{"page": "burn"})
}

func (sh *secretHandlers) ShowSecret(w http.ResponseWriter, r *http.Request) {
	utils.JsonResponse(w, map[string]string{"page": "show"})
}
