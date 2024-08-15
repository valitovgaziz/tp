package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func Test_hashPass(t *testing.T) {
	password := "some hard password"
	hash, _ := hashPassword(password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		t.Errorf("Falis by: %s", err)
	}
}
