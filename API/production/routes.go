package production

import (
	gr "db/API/graphic"
	"db/API/productionschedule"
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

	router.GET("/productionschedule", productionschedule.GetProductionschedule)
	router.GET("/productionschedule/:jobname", productionschedule.GetProductionscheduleByID)
	router.POST("/productionschedule", productionschedule.AddProductionschedule)
	router.PUT("/productionschedule/:jobname", productionschedule.UpdateProductionscheduleByID)
	router.DELETE("/productionschedule/:jobname", productionschedule.DeleteProductionscheduleByID)

	router.Run("localhost:8081")
	fmt.Printf("starting server at 8081")

}
