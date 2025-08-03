package handlers

import (
	"net/http"

	"github.com/alamgir-ahosain/go-rest-api-3ways/2_gorila_mux/internal/services"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	services.HandleCORSFunc(w)                //handle CORS error
	services.HandlePreflightRequestFunc(w, r) //handle Preflight request: for option method

	if r.Method != http.MethodGet {
		http.Error(w, "only GET request are allowed", 400)
		return
	}
	w.WriteHeader(200)
	services.MakeJSONFormatFunc(w, 200)

}

// POST request:create product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	services.HandleCORSFunc(w)
	services.HandlePreflightRequestFunc(w, r)
	if r.Method != http.MethodPost {
		http.Error(w, "only POST request are allowed", 400)
		return
	}
	services.CreateProductFunc(w, r)

}

// get aa product with Id
func GetProductWithID(w http.ResponseWriter, r *http.Request) {
	services.HandleCORSFunc(w)
	services.HandlePreflightRequestFunc(w, r)

	id, err := services.GetID(w, r)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	pr, err2 := services.GetProductWithIDFunc(id)
	if err2 != nil {
		http.Error(w, "id not found", http.StatusNotFound)
		return
	}
	services.JSON_with_threeParameterFunc(w, 201, pr)

}

// PUT request:partial update
func PatchProduct(w http.ResponseWriter, r *http.Request) {
	services.HandleCORSFunc(w)
	services.HandlePreflightRequestFunc(w, r)

	id, err := services.GetID(w, r)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	pr, err2 := services.PatchProductFunc(id)
	if err2 != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	services.JSON_with_threeParameterFunc(w, 200, pr)

}

// PUT request: full update or replce the entire object
func PutProduct(w http.ResponseWriter, r *http.Request) {
	services.HandleCORSFunc(w)
	services.HandlePreflightRequestFunc(w, r)
	id, err := services.GetID(w, r)
	if err != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	services.PutProductFunc(w, r, id)

}

// delete a specific product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {}
