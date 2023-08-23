package controller

import (
	"fmt"
	"db/dataservice"
	"db/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var graphics []models.Graphic

func GetGraphics(c *gin.Context) {
	fmt.Println("To get all the graphics present")
	graphics = dataservice.GetGraphicCardFromDb(nil)
	c.IndentedJSON(http.StatusOK, graphics)
}

func GetGraphicByID(c *gin.Context) {
	graphicrevisionnbr := c.Param("graphicrevisionnbr")
	fmt.Println(graphicrevisionnbr)
	//graphics = dataservice.GetGraphicCardFromDb(graphics)
	graphics = dataservice.GetGraphicCardFromDb(nil)

	// You're looping over graphics, but you haven't shown where graphics is defined.
	// Assuming graphics is a slice of models.Graphic.
	for _, a := range graphics {
		fmt.Println(a.GraphicRevisionNbr, graphicrevisionnbr)
		graphicrevisionnbr, _ := strconv.Atoi(graphicrevisionnbr)
		if a.GraphicRevisionNbr == graphicrevisionnbr {
			fmt.Println(a.GraphicRevisionNbr, graphicrevisionnbr)
			// Connect to the database and fetch the graphic card by graphicrevisionnbr
			graphicbygraphicrevisionnbr, err := dataservice.GetGraphicByGraphicrevisionnbrFromDb(graphicrevisionnbr)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve graphic card data"})
				return
			}
			// Return the fetched graphic card
			c.IndentedJSON(http.StatusOK, graphicbygraphicrevisionnbr)
			return
		}
	}

	// If no matching graphic card was found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Graphic not found"})
}

func AddGraphic(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var newgraphics models.Graphic

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newgraphics); err != nil {
		// fmt.Println("ERROR in getting data")
		return
	}
	// Add the new album to the slice.
	fmt.Println(newgraphics)
	//db := dataservice.ConnectToDb()
	dataservice.PostAddGraphicToDb(newgraphics)
	graphics = append(graphics, newgraphics)
	fmt.Println(graphics)
	c.IndentedJSON(http.StatusCreated, graphics)
}

func DeleteGraphicByID(c *gin.Context) {
	graphics = dataservice.GetGraphicCardFromDb(nil)
	graphicrevisionnbr := c.Param("graphicrevisionnbr")
	fmt.Println(graphicrevisionnbr)
	for index, a := range graphics {
		fmt.Println("In for id is:", a.GraphicRevisionNbr)
		graphicrevisionnbr, _ := strconv.Atoi(graphicrevisionnbr)
		if a.GraphicRevisionNbr == graphicrevisionnbr {
			//db := dataservice.ConnectToDb()
			dataservice.DeleteGraphicAlbumFromDb(a.GraphicRevisionNbr)
			graphics = append(graphics[:index], graphics[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "graphic not found"})

}

func UpdateGraphicByID(c *gin.Context) {
	graphics = dataservice.GetGraphicCardFromDb(nil)
	graphicrevisionnbr := c.Param("graphicrevisionnbr")
	fmt.Println(graphicrevisionnbr)
	// Parse the request body to get the updated album data
	var updatedGraphic models.Graphic
	if err := c.ShouldBindJSON(&updatedGraphic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, a := range graphics {
		graphicrevisionnbr, _ := strconv.Atoi(graphicrevisionnbr)
		if a.GraphicRevisionNbr == graphicrevisionnbr {
			fmt.Println("In for Graphicrevisionnbr is:", a.GraphicRevisionNbr)
			//db := dataservice.ConnectToDb()
			dataservice.UpdateGraphicInDb(updatedGraphic)
			// Update the album's data with the new data from the request body
			graphics[index].GraphicName = updatedGraphic.GraphicName
			graphics[index].GraphicRevisionNbr = updatedGraphic.GraphicRevisionNbr
			graphics[index].GraphicDesc = updatedGraphic.GraphicDesc
			graphics[index].GraphicFileName = updatedGraphic.GraphicFileName
			graphics[index].GraphicImage = updatedGraphic.GraphicImage
			graphics[index].DefaultTopPos = updatedGraphic.DefaultTopPos
			graphics[index].DefaultLeftPos = updatedGraphic.DefaultLeftPos
			graphics[index].DefaultColorCode = updatedGraphic.DefaultColorCode
			graphics[index].DefaultLocationCode = updatedGraphic.DefaultLocationCode
			graphics[index].DefaultHeight = updatedGraphic.DefaultHeight
			graphics[index].DefaultWidth = updatedGraphic.DefaultWidth
			graphics[index].EffectiveDate = updatedGraphic.EffectiveDate
			graphics[index].EndDate = updatedGraphic.EndDate
			graphics[index].LastUpdateTs = updatedGraphic.LastUpdateTs
			graphics[index].LastUpdateUserId = updatedGraphic.LastUpdateUserId
			graphics[index].VertOffset = updatedGraphic.VertOffset
			graphics[index].HorizOffset = updatedGraphic.HorizOffset

			c.IndentedJSON(http.StatusOK, graphics[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "graphics not found"})
}
