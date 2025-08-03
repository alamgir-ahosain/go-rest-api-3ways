package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alamgir-ahosain/go-rest-api-3ways/3_gin_Framework/internal/models"
	"github.com/alamgir-ahosain/go-rest-api-3ways/3_gin_Framework/internal/services"
	"github.com/gin-gonic/gin"
)

// GET request
func GetProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, services.List()) //code :200
}

// POST request
func AddProduct(c *gin.Context) {

	var newProduct models.Product
	err := c.BindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Eor:": err.Error()})
		return
	}
	services.Add(newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct) //code:201
}

// GET with id
func GetProduct(c *gin.Context) {

	fmt.Println("getProduct() called") // Add this line
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	pr, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error:": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, pr) //code :200
}

// PATCH request
func UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	pr, err := services.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error:": err.Error()})
		return
	}
	services.UpdateProduct(pr)
	c.IndentedJSON(http.StatusOK, pr) //code :200
}

// PUT request
func PutProduct(c *gin.Context) {
	idPam := c.Param("id")
	id, err := strconv.Atoi(idPam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}

	var newProduct models.Product
	err2 := c.BindJSON(&newProduct)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error:": err2.Error()})
		return
	}

	updateProduct, updated := services.PutProduct(id, &newProduct)
	if updated {
		c.JSON(http.StatusOK, updateProduct)
	} else {
		c.JSON(http.StatusCreated, updateProduct)
	}

}

// DELETE request
func DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)
	err := services.DeleteProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error:": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"product:": "Deleted Succesfully"}) //code :200
}
