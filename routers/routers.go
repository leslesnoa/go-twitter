package routers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-twitter/handlers"
	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/middleware"
	"github.com/rs/cors"
)

var (
	webURI = os.Getenv("WEB_URI")
)

func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(handlers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(handlers.Login)).Methods("POST")
	router.HandleFunc("/search", middleware.CheckDB(middleware.ValidJWT(handlers.SearchProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middleware.CheckDB(middleware.ValidJWT(handlers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/deleteUser", middleware.CheckDB(middleware.ValidJWT(handlers.DeleteUser))).Methods("DELETE")

	router.HandleFunc("/tweet", middleware.CheckDB(middleware.ValidJWT(handlers.PostTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middleware.CheckDB(middleware.ValidJWT(handlers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middleware.CheckDB(middleware.ValidJWT(handlers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/insertRelation", middleware.CheckDB(middleware.ValidJWT(handlers.Relation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middleware.CheckDB(middleware.ValidJWT(handlers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/consultRelation", middleware.CheckDB(middleware.ValidJWT(handlers.ConsultRelation))).Methods("GET")

	router.HandleFunc("/uploadAvatar", middleware.CheckDB(middleware.ValidJWT(handlers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middleware.CheckDB(handlers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middleware.CheckDB(middleware.ValidJWT(handlers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middleware.CheckDB(handlers.GetBanner)).Methods("GET")

	router.HandleFunc("/listUsers", middleware.CheckDB(middleware.ValidJWT(handlers.ListUsers))).Methods("GET")
	router.HandleFunc("/readFollowTweets", middleware.CheckDB(middleware.ValidJWT(handlers.ReadFollowTweets))).Methods("GET")

	if webURI == "" {
		webURI = "*"
		// webURI = "http://localhost:3000"
	}

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{webURI},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "OPTIONS", "DELETE", "POST", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		Debug:            true,
	}).Handler(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	logger.Info("about to start the application on PORT:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
