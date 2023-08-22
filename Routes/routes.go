package routes

import (
	"fmt"
	cr "go_API_CBCardDev/API/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	fmt.Println("In Routes")
	router.GET("/production", cr.GetProduction)
	router.GET("/production/:contractnbr", cr.GetProductionByID)
	router.POST("/production", cr.AddProduction)
	router.PUT("/production/:contractnbr", cr.UpdateProductionByID)
	router.DELETE("/production/:contractnbr", cr.DeleteProductionByID)

	router.GET("/graphic", cr.GetGraphics)
	router.GET("/graphic/:graphicrevisionnbr", cr.GetGraphicByID)
	router.POST("/graphic", cr.AddGraphic)
	router.PUT("/graphic/:graphicrevisionnbr", cr.UpdateGraphicByID)
	router.DELETE("/graphic/:graphicrevisionnbr", cr.DeleteGraphicByID)

	router.GET("/processedfile", cr.Getprocessedfiles)
	router.GET("/processedfile/:fileid", cr.GetprocessedfileByID)
	router.POST("/processedfile", cr.Addprocessedfile)
	router.PUT("/processedfile/:fileid", cr.UpdateprocessedfileByID)
	router.DELETE("/processedfile/:fileid", cr.DeleteprocessedfileByID)

	router.GET("/production", cr.GetCard)
	router.GET("/production/:cardrevisionid", cr.GetCardByID)
	router.POST("/production", cr.AddCard)
	router.PUT("/production/:cardrevisionid", cr.UpdateCardByID)
	router.DELETE("/production/:cardrevisionid", cr.DeleteCardByID)

	router.Run("localhost:8081")
	fmt.Printf("starting server at 8081")

}
