package controller

import (
	"fmt"
	"go_API_CBCardDev/dataservice"
	"go_API_CBCardDev/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// _ "github.com/gin-gonic/gin"
)

var platforms []models.Platform

func GetPlatforms(c *gin.Context) {
	fmt.Println("To get all the platforms present")
	// db := dataservice.ConnectToDb()
	// fmt.Println("\n in getPlatform", db)
	platforms = dataservice.GetPlatformFromDb(nil)
	c.IndentedJSON(http.StatusOK, platforms)
}

func GetPlatformByID(c *gin.Context) {
	idtype := c.Param("idtype")
	fmt.Println(idtype)
	//graphics = dataservice.GetGraphicCardFromDb(graphics)

	// You're looping over graphics, but you haven't shown where graphics is defined.
	// Assuming graphics is a slice of models.Graphic.
	for _, a := range platforms {
		fmt.Println(a.Idtype, idtype)
		//idtype, _ := strconv.Atoi(idtype)
		if a.Idtype == idtype {
			fmt.Println(a.Idtype, idtype)
			// Connect to the database and fetch the graphic card by graphicrevisionnbr
			idtype, _ := strconv.Atoi(idtype)
			platformbyidtype, err := dataservice.GetPlatformByIdtypeFromDb(idtype)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve platformcard data"})
				return
			}
			// Return the fetched graphic card
			c.IndentedJSON(http.StatusOK, platformbyidtype)
			return
		}
	}
	// If no matching graphic card was found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Platform not found"})
}

func AddPlatform(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var newplatforms models.Platform

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newplatforms); err != nil {
		// fmt.Println("ERROR in getting data")
		return
	}
	// Add the new album to the slice.
	fmt.Println(newplatforms)
	//db := dataservice.ConnectToDb()
	dataservice.PostAddPlatformToDb(newplatforms)
	platforms = append(platforms, newplatforms)
	fmt.Println(platforms)
	c.IndentedJSON(http.StatusCreated, platforms)
}

func DeletePlatformByID(c *gin.Context) {
	idtype := c.Param("idtype")
	fmt.Println(idtype)
	for index, a := range platforms {
		fmt.Println("In for id is:", a.Idtype)
		//idtype, _ := strconv.Atoi(idtype)
		if a.Idtype == idtype {
			//db := dataservice.ConnectToDb()
			dataservice.DeleteAlbumFromDb(a.Idtype)
			platforms = append(platforms[:index], platforms[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "platforms not found"})
}

func UpdatePlatformByID(c *gin.Context) {
	idtype := c.Param("idtype")
	fmt.Println(idtype)
	// Parse the request body to get the updated album data
	var updatedPlatform models.Platform
	if err := c.ShouldBindJSON(&updatedPlatform); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, a := range platforms {
		if a.Idtype == idtype {
			fmt.Println("In for  isIdtype:", a.Idtype)
			//db := dataservice.ConnectToDb()

			dataservice.UpdatePlatformInDb(updatedPlatform)
			// Update the album's data with the new data from the request body
			platforms[index].Idtype = updatedPlatform.Idtype
			platforms[index].Priority = updatedPlatform.Priority
			platforms[index].Description = updatedPlatform.Description
			platforms[index].Plasticstock = updatedPlatform.Plasticstock
			platforms[index].Carrierstock = updatedPlatform.Carrierstock
			platforms[index].Envelopestock = updatedPlatform.Envelopestock
			platforms[index].Setup = updatedPlatform.Setup
			platforms[index].Fmodule1 = updatedPlatform.Fmodule1
			platforms[index].Fmodule2 = updatedPlatform.Fmodule2
			platforms[index].Fmodule3 = updatedPlatform.Fmodule3
			platforms[index].Bmodule1 = updatedPlatform.Bmodule1
			platforms[index].Bmodule2 = updatedPlatform.Bmodule2
			platforms[index].Jobgroup = updatedPlatform.Jobgroup
			platforms[index].Defaultqueue = updatedPlatform.Defaultqueue
			platforms[index].Platformcode = updatedPlatform.Platformcode

			c.IndentedJSON(http.StatusOK, platforms[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "platforms not found"})
}
