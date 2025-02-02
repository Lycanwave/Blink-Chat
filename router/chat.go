package router

import (
	"go-lang/blinkchat/controllers"

	"github.com/gorilla/mux"
)

func ChatRouter(parentRouter *mux.Router) *mux.Router {

	router := parentRouter.PathPrefix("/api/chat").Subrouter()

	router.HandleFunc("/", controllers.CreateChat).Methods("POST")
	router.HandleFunc("/{id}", controllers.GetChat).Methods("GET")

	return router
}
