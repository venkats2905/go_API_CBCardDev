package controller

import (
	"fmt"
	"go_API_CBCardDev/dataservice"
	"go_API_CBCardDev/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var productioncards []models.Productioncard

func GetProduction(c *gin.Context) {
	fmt.Println("To get all the productioncards present")
	productioncards = dataservice.GetProductionCardFromDb(nil)
	c.IndentedJSON(http.StatusOK, productioncards)
}

func GetProductionByID(c *gin.Context) {
	contractnbr := c.Param("contractnbr")
	fmt.Println(contractnbr)

	// You're looping over productioncards, but you haven't shown where productioncards is defined.
	// Assuming productioncards is a slice of models.Productioncard.
	for _, a := range productioncards {
		fmt.Println(a.Contractnbr, contractnbr)
		if a.Contractnbr == contractnbr {
			fmt.Println(a.Contractnbr, contractnbr)
			// Connect to the database and fetch the production card by contractnbr
			productioncardbycontractnbr, err := dataservice.GetProductioncardByContractnbrFromDb(contractnbr)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve production card data"})
				return
			}
			// Return the fetched production card
			c.IndentedJSON(http.StatusOK, productioncardbycontractnbr)
			return
		}
	}

	// If no matching production card was found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Production card not found"})
}

func AddProduction(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var newproductioncards models.Productioncard

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newproductioncards); err != nil {
		// fmt.Println("ERROR in getting data")
		return
	}
	// Add the new album to the slice.
	fmt.Println(newproductioncards)
	//db := dataservice.ConnectToDb()
	dataservice.PostAddProductionToDb(newproductioncards)
	productioncards = append(productioncards, newproductioncards)
	fmt.Println(productioncards)
	c.IndentedJSON(http.StatusCreated, newproductioncards)
}

func DeleteProductionByID(c *gin.Context) {
	productioncards = dataservice.GetProductionCardFromDb(nil)
	contractnbr := c.Param("contractnbr")
	fmt.Println(contractnbr)
	for index, a := range productioncards {
		fmt.Println("In for id is:", a.Contractnbr)
		if a.Contractnbr == contractnbr {
			//db := dataservice.ConnectToDb()
			dataservice.DeleteAlbumFromDb(a.Contractnbr)
			productioncards = append(productioncards[:index], productioncards[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "productioncard not found"})

}

func UpdateProductionByID(c *gin.Context) {
	productioncards = dataservice.GetProductionCardFromDb(nil)
	contractnbr := c.Param("contractnbr")
	fmt.Println(contractnbr)
	// Parse the request body to get the updated album data
	var updatedProductioncard models.Productioncard
	if err := c.ShouldBindJSON(&updatedProductioncard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, a := range productioncards {
		if a.Contractnbr == contractnbr {
			fmt.Println("In for Contractnbr is:", a.Contractnbr)
			//db := dataservice.ConnectToDb()
			dataservice.UpdateProductionInDb(updatedProductioncard)
			// Update the album's data with the new data from the request body
			productioncards[index].Contractnbr = updatedProductioncard.Contractnbr
			productioncards[index].Requestdate = updatedProductioncard.Requestdate
			productioncards[index].Status = updatedProductioncard.Status
			productioncards[index].Statusdate = updatedProductioncard.Statusdate
			productioncards[index].Statusby = updatedProductioncard.Statusby
			productioncards[index].Searchcode = updatedProductioncard.Searchcode
			productioncards[index].Cardcount = updatedProductioncard.Cardcount
			productioncards[index].Jobname = updatedProductioncard.Jobname
			productioncards[index].Producedby = updatedProductioncard.Producedby
			productioncards[index].Produceddate = updatedProductioncard.Produceddate
			productioncards[index].Scheduleddate = updatedProductioncard.Scheduleddate
			productioncards[index].Cardtemplatecode = updatedProductioncard.Cardtemplatecode
			productioncards[index].Groupnbr = updatedProductioncard.Groupnbr
			productioncards[index].Suffixnbr = updatedProductioncard.Suffixnbr
			productioncards[index].Matrldist = updatedProductioncard.Matrldist
			productioncards[index].Trancd = updatedProductioncard.Trancd
			productioncards[index].Reasoncd = updatedProductioncard.Reasoncd
			productioncards[index].Reptype = updatedProductioncard.Reptype
			productioncards[index].Litcode = updatedProductioncard.Litcode

			c.IndentedJSON(http.StatusOK, productioncards[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "productioncards not found"})
}
