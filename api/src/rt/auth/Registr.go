package auth

import (
	"api/src/models"
	"api/src/storages/psql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var Crenetials models.Crenetials
	// Decoe body
	if err := json.NewDecoder(r.Body).Decode(&Crenetials); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// shep password
	hashedPassword, err := hashPassword(Crenetials.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id := uuid.New()

	user := models.User{
		Id:       id,
		Name:     Crenetials.Name,
		Email:    Crenetials.Email,
		Password: hashedPassword,
		Phone:    Crenetials.Phone,
	}
	if result := psql.PSQL_GORM_DB.Create(&user); result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
