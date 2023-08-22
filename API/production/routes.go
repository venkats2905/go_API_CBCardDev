package production

import (
	gr "db/API/graphic"

	"db/API/trackingsequence"

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

	router.GET("/graphic", gr.GetGraphics)
	router.GET("/graphic/:graphicrevisionnbr", gr.GetGraphicByID)
	router.POST("/graphic", gr.AddGraphic)
	router.PUT("/graphic/:graphicrevisionnbr", gr.UpdateGraphicByID)
	router.DELETE("/graphic/:graphicrevisionnbr", gr.DeleteGraphicByID)

	router.GET("/tracking", trackingsequence.GetTracking_Seq)
	router.GET("/tracking/:CONTRACTNBR", trackingsequence.GetTracking_SeqByID)
	router.POST("/tracking", trackingsequence.AddTracking_Seq)
	router.PUT("/tracking/:CONTRACTNBR", trackingsequence.UpdateTracking_SeqByID)
	router.DELETE("/tracking/:CONTRACTNBR", trackingsequence.DeleteTracking_SeqByID)

	router.Run("localhost:8081")
	fmt.Printf("starting server at 8081")

}
