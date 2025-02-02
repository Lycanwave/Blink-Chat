package router

import (
	"go-lang/blinkchat/controllers"

	"github.com/gorilla/mux"
)

func UserRouter(parentRouter *mux.Router) *mux.Router {

	router := parentRouter.PathPrefix("/api/user").Subrouter()

	router.HandleFunc("/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	router.HandleFunc("/", controllers.GetUsers).Methods("GET")

	return router
}
