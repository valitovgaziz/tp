package auth

import (
	"api/src/models"
	"api/src/storages/psql"
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check user
	var user models.User
	// get user by email
	result := psql.PSQL_GORM_DB.Where("email = ?", creds.Email).First(&user)
	if result.Error != nil || !checkPasswordHash(creds.Password, user.Password) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// create jwt token
	expirationtime := time.Now().Add(5 * time.Minute)
	claims := &models.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationtime),
		},
		Email: user.Email,
		Phone: user.Phone,
		Role:  user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationtime,
	})
	w.WriteHeader(http.StatusOK)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
