package db

import (
	"context"
	"time"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (models.UserInfo, bool) {
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	user, isExist, _ := CheckIsExistUser(email, ctx)
	if isExist == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		logger.Error("Error while CompareHashAndPassword process", err)
		return user, false
	}
	return user, true
}
