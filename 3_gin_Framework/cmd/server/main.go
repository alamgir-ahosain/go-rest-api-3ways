package main

import (
	"github.com/alamgir-ahosain/go-rest-api-3ways/3_gin_Framework/internal/product"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	product.RegisterRoutes(router)

	err := router.Run("localhost:8080")
	if err != nil {
		panic("Failed to start server" + err.Error())
	}
}
