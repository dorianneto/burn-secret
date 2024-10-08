package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/dorianneto/burn-secret/internal/interfaces"
	"github.com/dorianneto/burn-secret/internal/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authHandlers struct {
	database interfaces.KeyPairBased
	logger   *slog.Logger
}

func NewAuthHandlers(database interfaces.KeyPairBased, logger *slog.Logger) *authHandlers {
	return &authHandlers{database: database, logger: logger}
}

func (ah *authHandlers) Login(w http.ResponseWriter, r *http.Request) {
	type Input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input Input

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		ah.logger.Error("error on decoding input")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	id, err := ah.database.Get(fmt.Sprintf("user_email:%s", input.Email))
	if err != nil {
		ah.logger.Error("error on fetching user_email index")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	password, err := ah.database.Select(fmt.Sprintf("user:%s", id), "password")
	if err != nil {
		ah.logger.Error("error on fetching user")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(input.Password))
	if err != nil {
		ah.logger.Error("error on comparing password")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Create JWT token
	claims := jwt.MapClaims{
		"user":    input.Email,
		"expires": time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims, nil)

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		ah.logger.Error("error on creating token")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	utils.JsonResponse(w, map[string]string{"data": tokenString})
}
