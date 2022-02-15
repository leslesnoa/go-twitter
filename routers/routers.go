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

	noAuthRouter := router.MatcherFunc(IsNoAuthRouter).Subrouter()
	noAuthRouter.Use(middleware.CheckDB)

	noAuthRouter.HandleFunc("/register", handlers.Register).Methods("POST")
	noAuthRouter.HandleFunc("/login", handlers.Login).Methods("POST")
	noAuthRouter.HandleFunc("/getAvatar", handlers.GetAvatar).Methods("GET")
	noAuthRouter.HandleFunc("/getBanner", handlers.GetBanner).Methods("GET")

	authRouter := router.MatcherFunc(IsAuthRouter).Subrouter()
	authRouter.Use(middleware.CheckDB, middleware.ValidJWT)

	authRouter.HandleFunc("/search", handlers.SearchProfile).Methods("GET")
	authRouter.HandleFunc("/modifyProfile", handlers.ModifyProfile).Methods("PUT")
	authRouter.HandleFunc("/deleteUser", handlers.DeleteUser).Methods("DELETE")

	authRouter.HandleFunc("/tweet", handlers.PostTweet).Methods("POST")
	authRouter.HandleFunc("/readTweets", handlers.ReadTweets).Methods("GET")
	authRouter.HandleFunc("/deleteTweet", handlers.DeleteTweet).Methods("DELETE")

	authRouter.HandleFunc("/insertRelation", handlers.Relation).Methods("POST")
	authRouter.HandleFunc("/deleteRelation", handlers.DeleteRelation).Methods("DELETE")
	authRouter.HandleFunc("/consultRelation", handlers.ConsultRelation).Methods("GET")

	authRouter.HandleFunc("/listUsers", handlers.ListUsers).Methods("GET")
	authRouter.HandleFunc("/readFollowTweets", handlers.ReadFollowTweets).Methods("GET")

	authRouter.HandleFunc("/uploadAvatar", handlers.UploadAvatar).Methods("POST")
	authRouter.HandleFunc("/uploadBanner", handlers.UploadBanner).Methods("POST")

	if webURI == "" {
		webURI = "http://localhost:3000"
	}

	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{webURI},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "OPTIONS", "DELETE", "POST", "PUT"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		// Debug:            true,
	}).Handler(router)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	logger.Info("about to start the application on PORT:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func IsNoAuthRouter(r *http.Request, rm *mux.RouteMatch) bool {
	return r.Header.Get("Authorization") == ""
}

func IsAuthRouter(r *http.Request, rm *mux.RouteMatch) bool {
	return true
}
