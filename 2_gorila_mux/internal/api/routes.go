package api

import (
	"github.com/alamgir-ahosain/go-rest-api-3ways/2_gorila_mux/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products",handlers.GetProducts)
}