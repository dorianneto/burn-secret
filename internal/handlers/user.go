package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dorianneto/burn-secret/internal/interfaces"
	"github.com/dorianneto/burn-secret/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userHandlers struct {
	database interfaces.KeyPairBased
	logger   *slog.Logger
}

func NewUserHandlers(database interfaces.KeyPairBased, logger *slog.Logger) *userHandlers {
	return &userHandlers{database: database, logger: logger}
}

func (sh *userHandlers) RegisterUser(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Id       string `json:"id"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		sh.logger.Error("error on decoding input")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		sh.logger.Error("error on hashing password")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id := uuid.New().String()
	input.Id = id
	input.Password = string(passwordHashed)

	err = sh.database.Set(fmt.Sprintf("user_email:%s", input.Email), input.Id)
	if err != nil {
		sh.logger.Error("error on creating user_email index", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = sh.database.Insert(fmt.Sprintf("user:%s", id), map[string]interface{}{
		"id":       input.Id,
		"email":    input.Email,
		"password": input.Password,
	})
	if err != nil {
		sh.logger.Error("error on registering user", "error", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, map[string]string{"data": input.Id})
}
