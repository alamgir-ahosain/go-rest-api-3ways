package services

import (
	"encoding/json"
	"net/http"

	"github.com/alamgir-ahosain/go-rest-api-3ways/2_gorila_mux/internal/models"
)

var products = []models.Product{
	{ID: 1, Name: "Apple", Price: 10},
	{ID: 2, Name: "Mango", Price: 20},
	{ID: 3, Name: "Jack Fruit", Price: 100},
}

// Handle CORS error
func HandleCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                  //set header: (*) anyone can access th
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Alamgir-CustomHeader") //allow content type
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS") //allow content type
	w.Header().Set("Content-Type", "application/json")
}

// Handle Preflight Request for option method
func HandlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "only GET request are allowed", 400)
		return
	}
}

func GetProduct(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}
