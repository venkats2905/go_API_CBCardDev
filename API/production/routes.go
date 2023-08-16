package production

import (
	gr "db/API/graphic"
	tr "db/API/trackler"
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

	router.GET("/tracking", tr.GetProduction)
	router.GET("/tracking/:SID", tr.GetBySID)
	router.POST("/tracking", tr.Adddata)
	router.PUT("/tracking/:SID", tr.UpdateProductionByID)
	router.DELETE("/tracking/:SID", tr.DeleteProductionByID)

	router.Run("localhost:8081")
	fmt.Printf("starting server at 8081")

}
