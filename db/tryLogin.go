package db

import (
	"context"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string, ctx context.Context) (*models.UserInfo, bool) {

	user, isExist, _ := CheckIsExistUser(email, ctx)
	if !isExist {
		return nil, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		logger.Error("Error while CompareHashAndPassword process", err)
		return nil, false
	}
	return user, true
}
