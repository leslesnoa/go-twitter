package main

import (
	"log"

	"github.com/leslesnoa/go-twitter/db"
	"github.com/leslesnoa/go-twitter/handlers"
)

func main() {
	if db.CheckingConnection() == 0 {
		log.Fatal("db connection error")
		return
	}
	log.Println("DB connection success")
	handlers.Handler()
}
