package helpers

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func ValidateMail(m string) bool {
	_, err := mail.ParseAddress(m)

	return err == nil
}
