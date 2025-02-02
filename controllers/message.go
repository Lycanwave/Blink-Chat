package controllers

import (
	"encoding/json"
	"fmt"
	"go-lang/blinkchat/models"
	"go-lang/blinkchat/services"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {

	var request *models.Message

	err := json.NewDecoder(r.Body).Decode(&request)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	res, err := services.CreateMessage(request)

	if err != nil {
		http.Error(w, "Error creating Message", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	res, err := services.GetMessage(id)

	if err != nil {
		http.Error(w, "error in Getting this Message", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetMessages(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	res, err := services.GetMessages(id)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "error in Getting this Message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
