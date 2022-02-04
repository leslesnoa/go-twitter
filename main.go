package main

import (
	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/routers"
)

func main() {
	if err := db.CheckingConnection(); err != nil {
		logger.Error("DB connection error", err)
		return
	}
	logger.Info("DB connection success")
	routers.Router()
}
