package db

import (
	"context"
	"time"

	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

func TryLogin(email string, password string) (*models.UserInfo, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

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
