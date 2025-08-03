package api

import (
	"github.com/alamgir-ahosain/go-rest-api-3ways/2_gorila_mux/internal/handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/products", handlers.GetProducts).Methods("GET")           //get all product
	r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")        //create a new product
	r.HandleFunc("/products/{id}", handlers.GetProductWithID).Methods("GET") //get a specific product
	r.HandleFunc("/products/{id}", handlers.PatchProduct).Methods("PATCH")   //partial update
	r.HandleFunc("/products/{id}", handlers.PutProduct).Methods("PUT")       //full update or replace the entire object
	r.HandleFunc("/products/{id}", handlers.DeleteProduct).Methods("DELETE") //delete a product
}
