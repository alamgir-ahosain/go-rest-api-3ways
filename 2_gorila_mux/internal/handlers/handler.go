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
