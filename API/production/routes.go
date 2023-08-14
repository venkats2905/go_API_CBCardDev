package production

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	fmt.Println("In Routes")
	router.GET("/production", getProduction)
	router.GET("/production/:contractnbr", getProductionByID)
	router.POST("/production", addProduction)
	router.PUT("/production/:contractnbr", updateProductionByID)
	router.DELETE("/production/:contractnbr", deleteProductionByID)

	router.Run("localhost:8081")
	fmt.Printf("starting server at 8081")

}
