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
	// websocket
	router.Handle("/api/ws", http.HandlerFunc(get_ws)).Methods("GET")
	router.Handle("/api/chatRoom/{chatRoomId}", http.HandlerFunc(get_chatRoom_ws)).Methods("GET")

	// auth routes
	router.Handle("/api/jwt", http.HandlerFunc(get_jwt)).Methods("GET")
	router.Handle("/api/logout", http.HandlerFunc(get_logOut)).Methods("POST")
	router.Handle("/api/signup", http.HandlerFunc(post_signup)).Methods("POST")
	router.Handle("/api/login", http.HandlerFunc(post_login)).Methods("POST")
	router.Handle("/api/set-profile", http.HandlerFunc(post_set_profile)).Methods("POST")
	router.Handle("/api/set-profileImg", http.HandlerFunc(post_set_profile_img)).Methods("POST")
	router.Handle("/api/set-resume", http.HandlerFunc(post_set_resume)).Methods("POST")
	router.Handle("/api/reset-password", http.HandlerFunc(post_reset_password)).Methods("POST")
	// post routes
	router.Handle("/api/availablePosts", http.HandlerFunc(get_available_posts)).Methods("GET")
	router.Handle("/api/new-liked", http.HandlerFunc(get_new_likes)).Methods("GET")
	router.Handle("/api/post/{postId}", http.HandlerFunc(get_single_post)).Methods("GET")
	router.Handle("/api/liked-by/{postId}", http.HandlerFunc(get_liked_by)).Methods("GET")
	router.Handle("/api/liked-posts", http.HandlerFunc(get_liked_posts)).Methods("GET")
	router.Handle("/api/create-post", http.HandlerFunc(post_create_post)).Methods("POST")
	router.Handle("/api/create-postImg/{postId}", http.HandlerFunc(post_create_post_img)).Methods("POST")
	router.Handle("/api/like-post", http.HandlerFunc(post_like_post)).Methods("POST")
	router.Handle("/api/update-post", http.HandlerFunc(post_update_post)).Methods("POST")
	router.Handle("/api/delete-post", http.HandlerFunc(post_delete_post)).Methods("POST")
	// profile
	router.Handle("/api/profile/{username}", http.HandlerFunc(get_profile)).Methods("GET")
	router.Handle("/api/get-notifications", http.HandlerFunc(get_get_notification)).Methods("GET")
	// comments
	router.Handle("/api/add-comment", http.HandlerFunc(post_create_comment)).Methods("POST")
	router.Handle("/api/delete-Comment", http.HandlerFunc(post_delete_comment)).Methods("POST")
	// Chat
	router.Handle("/api/load-chatRooms", http.HandlerFunc(get_load_chatRooms)).Methods("GET")
	router.Handle("/api/get-chat", http.HandlerFunc(post_get_chat)).Methods("POST")
	router.Handle("/api/create-room", http.HandlerFunc(post_create_room)).Methods("POST")
	router.Handle("/api/create-message", http.HandlerFunc(post_create_message)).Methods("POST")
	router.Handle("/api/create-message-img/{TargetId}", http.HandlerFunc(post_create_message_img)).Methods("POST")
	router.Handle("/api/update-message", http.HandlerFunc(post_update_message)).Methods("POST")
	// Setting
	router.Handle("/api/setting", http.HandlerFunc(get_setting)).Methods("GET")
	router.Handle("/api/remove-profileImg", http.HandlerFunc(post_remove_profileImg)).Methods("POST")
	router.Handle("/api/update-profileImg", http.HandlerFunc(post_update_profileImg)).Methods("POST")
	router.Handle("/api/update-personal", http.HandlerFunc(post_update_personal)).Methods("POST")
	router.Handle("/api/update-account", http.HandlerFunc(post_update_account)).Methods("POST")
	router.Handle("/api/change-password", http.HandlerFunc(post_update_password)).Methods("POST")
	router.Handle("/api/update-professional", http.HandlerFunc(post_update_professional)).Methods("POST")
	router.Handle("/api/delete-account", http.HandlerFunc(post_delete_account)).Methods("POST")

	// friendship
	router.Handle("/api/friend-requests", http.HandlerFunc(get_friend_requests)).Methods("GET")
	router.Handle("/api/online-friends", http.HandlerFunc(get_online_friends)).Methods("GET")
	router.Handle("/api/friends/{username}", http.HandlerFunc(get_friends_by_username)).Methods("GET")
	router.Handle("/api/add-friend", http.HandlerFunc(post_add_friend)).Methods("POST")
	router.Handle("/api/accept-request", http.HandlerFunc(post_accept_request)).Methods("POST")
	router.Handle("/api/reject-request", http.HandlerFunc(post_reject_request)).Methods("POST")

	// Job
	router.Handle("/api/my-jobs", http.HandlerFunc(get_my_jobs)).Methods("GET")
	router.Handle("/api/find-jobs", http.HandlerFunc(get_find_jobs)).Methods("GET")
	router.Handle("/api/create-job", http.HandlerFunc(post_create_job)).Methods("POST")
	// search
	router.Handle("/api/search/{userInput}", http.HandlerFunc(get_search)).Methods("GET")
}
