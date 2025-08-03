package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/alamgir-ahosain/go-rest-api-3ways/2_gorila_mux/internal/models"
	"github.com/gorilla/mux"
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

// make JSON format
func MakeJSONFormatFunc(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(products)
}

// make JSON format
func JSON_with_threeParameterFunc(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

// get id
func GetID(w http.ResponseWriter, r *http.Request) (int, error) {
	ids := mux.Vars(r)
	idStr := ids["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		//http.Error(w, "invalid product id", http.StatusBadRequest)
		return 0, fmt.Errorf("invalid product id")
	}
	return id, nil
}

// create product function:post request
func CreateProductFunc(w http.ResponseWriter, r *http.Request) {
	var newProduct models.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give  the valid json format", 400)
		return
	}

	products = append(products, newProduct)
	MakeJSONFormatFunc(w, 201)
}

// get product with id function
func GetProductWithIDFunc(id int) (models.Product, error) {
	for i, val := range products {
		if val.ID == id {
			return products[i], nil
		}
	}
	return models.Product{}, fmt.Errorf("product with ID %d not found", id)
}

// Patch request: partial update
func PatchProductFunc(id int) (models.Product, error) {
	for i, val := range products {
		if val.ID == id {
			products[i].Price += 2
			return products[i], nil
		}
	}
	return models.Product{}, fmt.Errorf("product with ID %d not found", id)
}

// PUT product function
func PutProductFunc(w http.ResponseWriter, r *http.Request, id int) {
	var newProduct models.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Give  the valid json format", 400)
		return
	}

	for i, val := range products {
		if val.ID == id {
			val.ID = id
			products[i] = newProduct
			JSON_with_threeParameterFunc(w, 201, newProduct)
			return

		}
	}

	JSON_with_threeParameterFunc(w, 201, newProduct)
}

// delete specific product
func DeleteProductFunc(w http.ResponseWriter,r *http.Request,id int) {
	for i, val := range products {
		if val.ID == id {
			JSON_with_threeParameterFunc(w, 201, products[i])
			products = append(products[:i], products[i+1:]...)
		}
	}

}
