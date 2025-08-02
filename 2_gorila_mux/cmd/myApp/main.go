package main

import (
	"fmt"
	"net/http"
	"github.com/alamgir-ahosain/go-rest-api-3ways/2_gorila_mux/internal/api"
	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter() //create router
	api.RegisterRoutes(mux)
	fmt.Println("Server running on port :8080")
	err:=http.ListenAndServe(":8080", mux)
	if err!=nil{
		fmt.Println("Error starting the server")
	}


}
