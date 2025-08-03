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
func HandleCORSFunc(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")                                  //set header: (*) anyone can access th
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Alamgir-CustomHeader") //allow content type
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS") //allow content type
	w.Header().Set("Content-Type", "application/json")
}

// Handle Preflight Request for option method
func HandlePreflightRequestFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	
}

//make JSON format
func MakeJSONFormatFunc(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}


func CreateProductFunc(w http.ResponseWriter,r *http.Request){
	var newProduct models.Product
	decoder:=json.NewDecoder(r.Body)
	err:=decoder.Decode(&newProduct)
	if err!=nil{
		http.Error(w,"Give  the valid json format",400)
		return
	}

	products=append(products, newProduct)
	MakeJSONFormatFunc(w,201)
}