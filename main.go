package main

import (
	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/handlers"
	"github.com/leslesnoa/go-twitter/logger"
)

func main() {
	if err := db.CheckingConnection(); err != nil {
		logger.Error("DB connection error", err)
		return
	}
	logger.Info("DB connection success")
	handlers.Handler()
}
