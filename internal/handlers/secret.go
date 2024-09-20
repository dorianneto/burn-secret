package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dorianneto/burn-secret/internal/interfaces"
	"github.com/dorianneto/burn-secret/internal/utils"
	"github.com/google/uuid"
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

	cipherData, err := utils.EncryptIt(input.Secret)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id := uuid.New().String()

	sh.database.Insert(fmt.Sprintf("_:secret:%s", id), map[string]interface{}{
		"secret": cipherData.Code,
		"nonce":  cipherData.Nonce,
	})

	utils.JsonResponse(w, map[string]string{"data": id})
}

func (sh *secretHandlers) BurnSecret(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	deleted, err := sh.database.Delete(fmt.Sprintf("_:secret:%s", id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if deleted == 0 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, map[string]int64{"deleted": deleted})
}

func (sh *secretHandlers) ShowSecret(w http.ResponseWriter, r *http.Request) {
	type Secret struct {
		Code  string `redis:"secret"`
		Nonce string `redis:"nonce"`
	}

	var secret Secret

	id := r.PathValue("id")

	err := sh.database.SelectAll(fmt.Sprintf("_:secret:%s", id), &secret)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	secretParsed, err := utils.DecryptIt(utils.NewCipherData(secret.Code, []byte(secret.Nonce)))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, map[string]string{"data": secretParsed})
}
