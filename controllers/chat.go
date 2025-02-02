package controllers

import (
	"encoding/json"
	"go-lang/blinkchat/models"
	"go-lang/blinkchat/services"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {

	var request models.Chat

	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := services.CreateChat(request)

	if err != nil {
		http.Error(w, "Error creating chat", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetChat(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	res, err := services.GetChat(id)

	if err != nil {
		http.Error(w, "error in Getting this Chat", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
