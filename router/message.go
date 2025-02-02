package router

import (
	"go-lang/blinkchat/controllers"

	"github.com/gorilla/mux"
)

func MessageRouter(parentRouter *mux.Router) *mux.Router {

	router := parentRouter.PathPrefix("/api/message").Subrouter()

	router.HandleFunc("/", controllers.CreateMessage).Methods("POST")
	router.HandleFunc("/{id}", controllers.GetMessage).Methods("GET")
	router.HandleFunc("/{id}/chat", controllers.GetMessages).Methods("GET")

	return router
}
