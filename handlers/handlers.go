package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-twitter/middleware"
	"github.com/leslesnoa/go-twitter/routers"
	"github.com/rs/cors"
)

func Handler() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/search", middleware.ValidJWT(routers.SearchProfile)).Methods("GET")
	router.HandleFunc("/modifyProfile", middleware.ValidJWT(routers.ModifyProfile)).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
