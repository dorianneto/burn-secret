package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dorianneto/burn-secret/internal/interfaces"
	"github.com/dorianneto/burn-secret/internal/utils"
)

type secretHandlers struct {
	database interfaces.KeyPairBased
}

func NewSecretHandlers(database interfaces.KeyPairBased) *secretHandlers {
	return &secretHandlers{database: database}
}

func (sh *secretHandlers) GenerateSecret(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Secret string `json:"secret"`
	}

	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	utils.JsonResponse(w, input)
}

func (sh *secretHandlers) BurnSecret(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	utils.JsonResponse(w, map[string]string{"page": id})
}

func (sh *secretHandlers) ShowSecret(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	utils.JsonResponse(w, map[string]string{"page": id})
}
