package router

import (
	"go-lang/blinkchat/socket"

	"github.com/gorilla/mux"
)

func SocketRouter(parentRouter *mux.Router) *mux.Router {

	router := parentRouter.PathPrefix("/ws/api").Subrouter()
	router.HandleFunc("/message", socket.HandleWebSocket)

	return router
}
