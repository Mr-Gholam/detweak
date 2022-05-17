package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routerSetup(router *mux.Router) {

	router.Use(corsMiddleware)
	router.Use(authMiddleware)
	router.Handle("/api/jwt", http.HandlerFunc(get_jwt)).Methods("GET")
	router.Handle("/api/logout", http.HandlerFunc(get_logOut)).Methods("POST")
	router.Handle("/api/signup", http.HandlerFunc(post_signup)).Methods("POST")
	router.Handle("/api/login", http.HandlerFunc(post_login)).Methods("POST")
	router.Handle("/api/set-profile", http.HandlerFunc(post_set_profile)).Methods("POST")
	router.Handle("/api/set-resume", http.HandlerFunc(post_set_resume)).Methods("POST")
}
