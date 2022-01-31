package db

import (
	"github.com/leslesnoa/go-twitter/logger"
	"golang.org/x/crypto/bcrypt"
)

func EncriptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	if err != nil {
		logger.Error("Error while Encripting password", err)
	}
	return string(bytes), err
}
