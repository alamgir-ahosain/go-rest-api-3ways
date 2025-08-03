package api

import (
	
	"github.com/alamgir-ahosain/go-rest-api-3ways/3_gin_Framework/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/products", handlers.GetProducts)
	r.POST("/products",handlers.AddProduct) //create new one
	r.GET("/products/:id", handlers.GetProduct)
	r.PATCH("/products/:id", handlers.UpdateProduct) //partial update
	r.PUT("/products/:id", handlers.PutProduct)      //full update or replace the entire object
	r.DELETE("/products/:id", handlers.DeleteProduct)

}