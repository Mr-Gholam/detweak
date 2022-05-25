package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func routerSetup(router *mux.Router) {

	router.Use(corsMiddleware)
	router.Use(authMiddleware)
	// img handler
	router.Handle("/api/images/{imgUrl}", http.HandlerFunc(sendFileToClient)).Methods("GET")
	// auth routes
	router.Handle("/api/jwt", http.HandlerFunc(get_jwt)).Methods("GET")
	router.Handle("/api/logout", http.HandlerFunc(get_logOut)).Methods("POST")
	router.Handle("/api/signup", http.HandlerFunc(post_signup)).Methods("POST")
	router.Handle("/api/login", http.HandlerFunc(post_login)).Methods("POST")
	router.Handle("/api/set-profile", http.HandlerFunc(post_set_profile)).Methods("POST")
	router.Handle("/api/set-profileImg", http.HandlerFunc(post_set_profile_img)).Methods("POST")
	router.Handle("/api/set-resume", http.HandlerFunc(post_set_resume)).Methods("POST")
	// post routes
	router.Handle("/api/availablePosts", http.HandlerFunc(get_available_posts)).Methods("GET")
	router.Handle("/api/post/{postId}", http.HandlerFunc(get_single_post)).Methods("GET")
	router.Handle("/api/create-post", http.HandlerFunc(post_create_post)).Methods("POST")
	router.Handle("/api/create-postImg/{postId}", http.HandlerFunc(post_create_post_img)).Methods("POST")
	router.Handle("/api/like-post", http.HandlerFunc(post_like_post)).Methods("POST")
	router.Handle("/api/update-post", http.HandlerFunc(post_update_post)).Methods("POST")
	router.Handle("/api/delete-post", http.HandlerFunc(post_delete_post)).Methods("POST")
	// profile
	router.Handle("/api/profile/{username}", http.HandlerFunc(get_profile)).Methods("GET")
	// comments
	router.Handle("/api/add-comment", http.HandlerFunc(post_create_comment)).Methods("POST")
	router.Handle("/api/delete-Comment", http.HandlerFunc(post_delete_comment)).Methods("POST")

	// friendship
	router.Handle("/api/friend-requests", http.HandlerFunc(get_friend_requests)).Methods("GET")
	router.Handle("/api/add-friend", http.HandlerFunc(post_add_friend)).Methods("POST")
	router.Handle("/api/accept-request", http.HandlerFunc(post_accept_request)).Methods("POST")
	router.Handle("/api/reject-request", http.HandlerFunc(post_reject_request)).Methods("POST")

	// search
	router.Handle("/api/search/{userInput}", http.HandlerFunc(get_search)).Methods("GET")
}
