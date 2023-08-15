package trackler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/tracking", getProduction)
	router.GET("/tracking/:SID", getBySID)
	router.POST("/tracking", adddata)
	router.PUT("/tracking/:SID", updateProductionByID)
	router.DELETE("/tracking/:SID", deleteProductionByID)

	router.Run("localhost:8082")
	fmt.Printf("starting server at 8082")

}
