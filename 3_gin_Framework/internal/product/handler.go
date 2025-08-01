package product

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func RegisterRoutes(r * gin.Engine){
		r.GET("/products",getProducts)

}



func getProducts(c *gin.Context){
 c.IndentedJSON(http.StatusOK,List()) //code :200

}
