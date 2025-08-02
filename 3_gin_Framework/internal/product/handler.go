package product

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/products", getProducts)
	r.POST("/products", addProduct) //create new one
	r.GET("/products/:id", getProduct)
	r.PATCH("/products/:id", updateProduct) //partial update
	r.PUT("/products/:id", putProduct)   //full update or replace the entire object
	r.DELETE("/products/:id", deleteProduct)

}

// GET request
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, List()) //code :200

}

// GET with id
func getProduct(c *gin.Context) {

	fmt.Println("getProduct() called") // Add this line
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	pr, err := GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error:": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, pr) //code :200

}

// PATCH request
func updateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	pr, err := GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error:": err.Error()})
		return
	}
	UpdateProduct(pr)
	c.IndentedJSON(http.StatusOK, pr) //code :200
}



// DELETE request
func deleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	err := DeleteProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error:": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"product:": "Deleted Succesfully"}) //code :200
}

// POST request
func addProduct(c *gin.Context) {

	var newProduct Product
	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Eor:": err.Error()})
		return
	}
	Add(newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct) //code:201

}
