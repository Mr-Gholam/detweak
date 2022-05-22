package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//make connection to database
	connectDb()
	// setup the router
	router := mux.NewRouter()

	routerSetup(router)
	http.ListenAndServe("127.0.0.1:8585", router)
}
