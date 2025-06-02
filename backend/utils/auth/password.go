package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(raw_password string) (string, error) {

	hashed, err := bcrypt.GenerateFromPassword([]byte(raw_password), bcrypt.DefaultCost)

	return string(hashed), err

}

func CheckPassword(raw_password string, hashed_password string) bool {

	// fmt.Printf("Raw Password: %s; Hashed Password: %s", raw_password, hashed_password)

	err := bcrypt.CompareHashAndPassword([]byte(hashed_password), []byte(raw_password))

	return err == nil
}
