package production

import (
	gr "db/API/graphic"
	pt "db/API/platform"
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

	router.GET("/platform", pt.GetPlatforms)
	router.GET("/platform/:idtype", pt.GetPlatformByID)
	router.POST("/platform", pt.AddPlatform)
	router.PUT("/platform/:idtype", pt.UpdatePlatformByID)
	router.DELETE("/platform/:idtype", pt.DeletePlatformByID)

	router.Run("localhost:8081")
	fmt.Printf("starting server at 8081")

}
