package router

import "github.com/gorilla/mux"

func Router() *mux.Router {

	router := mux.NewRouter()
	router.PathPrefix("/ws/api").Handler(SocketRouter(router))
	router.PathPrefix("/api/user").Handler(UserRouter(router))
	router.PathPrefix("/api/chat").Handler(ChatRouter(router))
	router.PathPrefix("/api/message").Handler(MessageRouter(router))

	return router
}
