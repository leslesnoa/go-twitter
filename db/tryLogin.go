package db

import (
	"github.com/leslesnoa/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.UserInfo, bool) {
	user, isExist, _ := CheckIsExistUser(email)
	if isExist == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
