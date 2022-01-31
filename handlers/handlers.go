package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leslesnoa/go-twitter/logger"
	"github.com/leslesnoa/go-twitter/middleware"
	"github.com/leslesnoa/go-twitter/routers"
	"github.com/rs/cors"
)

func Handler() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middleware.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middleware.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/search", middleware.CheckDB(middleware.ValidJWT(routers.SearchProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middleware.CheckDB(middleware.ValidJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.CheckDB(middleware.ValidJWT(routers.PostTweet))).Methods("POST")
	router.HandleFunc("/readTweets", middleware.CheckDB(middleware.ValidJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deleteTweet", middleware.CheckDB(middleware.ValidJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/insertRelation", middleware.CheckDB(middleware.ValidJWT(routers.Relation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middleware.CheckDB(middleware.ValidJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/consultRelation", middleware.CheckDB(middleware.ValidJWT(routers.ConsultRelation))).Methods("GET")

	router.HandleFunc("/uploadAvatar", middleware.CheckDB(middleware.ValidJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middleware.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadBanner", middleware.CheckDB(middleware.ValidJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middleware.CheckDB(routers.GetBanner)).Methods("GET")

	router.HandleFunc("/listUsers", middleware.CheckDB(middleware.ValidJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/readFollowTweets", middleware.CheckDB(middleware.ValidJWT(routers.ReadFollowTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)

	logger.Info("about to start the application...")
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
