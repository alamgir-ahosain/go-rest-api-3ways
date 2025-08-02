package handlers

import (
	"net/http"

	"github.com/alamgir-ahosain/go-rest-api-3ways/2_gorila_mux/internal/services"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {

	services.HandleCORS(w)                //handle CORS error
	services.HandlePreflightRequest(w, r) //handle Preflight request: for option method
	w.WriteHeader(200)
	services.GetProduct(w, 200)

}
