package trackingsequence

import (
	"db/dataservice"
	"db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// _ "github.com/gin-gonic/gin"
)

var trackingtemp []models.Tracking_seq

func GetTracking_Seq(c *gin.Context) {
	fmt.Println("To get all the trackingtemp present")
	// db := dataservice.ConnectToDb()
	// fmt.Println("\n in getProduction", db)
	trackingtemp = dataservice.GetTracking_SeqFromDb(nil)
	c.IndentedJSON(http.StatusOK, trackingtemp)
	//trackingtemp = make([]models.Productioncard, 0)
}

func GetTracking_SeqByID(c *gin.Context) {
	CONTRACTNBR := c.Param("CONTRACTNBR")
	fmt.Println("contractnbr", CONTRACTNBR)
	trackingtemp = dataservice.GetTracking_SeqFromDb(nil)

	// You're looping over trackingtemp, but you haven't shown where trackingtemp is defined.
	// Assuming trackingtemp is a slice of models.Productioncard.
	for _, a := range trackingtemp {
		fmt.Println(a.CONTRACTNBR, CONTRACTNBR)
		if a.CONTRACTNBR == CONTRACTNBR {
			fmt.Println(a.CONTRACTNBR, CONTRACTNBR)
			// Connect to the database and fetch the production card by CONTRACTNBR
			productioncardbyCONTRACTNBR, err := dataservice.GetTracking_SeqByContractnbrFromDb(CONTRACTNBR)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve production card data"})
				return
			}
			// Return the fetched production card
			c.IndentedJSON(http.StatusOK, productioncardbyCONTRACTNBR)
			return
		}
	}

	// If no matching production card was found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Production card not found"})
}

func AddTracking_Seq(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var newtrackingtemp models.Tracking_seq

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newtrackingtemp); err != nil {
		// fmt.Println("ERROR in getting data")
		return
	}
	// Add the new album to the slice.
	fmt.Println(newtrackingtemp)
	//db := dataservice.ConnectToDb()
	dataservice.PostAddTracking_SeqAlbumToDb(newtrackingtemp)
	trackingtemp = append(trackingtemp, newtrackingtemp)
	fmt.Println(trackingtemp)
	c.IndentedJSON(http.StatusCreated, newtrackingtemp)
}

func DeleteTracking_SeqByID(c *gin.Context) {
	CONTRACTNBR := c.Param("CONTRACTNBR")
	fmt.Println(CONTRACTNBR)
	trackingtemp = dataservice.GetTracking_SeqFromDb(nil)
	for index, a := range trackingtemp {
		fmt.Println("In for id is:", a.CONTRACTNBR)
		if a.CONTRACTNBR == CONTRACTNBR {
			//db := dataservice.ConnectToDb()
			dataservice.DeleteTracking_SeqAlbumFromDb(a.CONTRACTNBR)
			trackingtemp = append(trackingtemp[:index], trackingtemp[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "trackingsequence not found"})

}

func UpdateTracking_SeqByID(c *gin.Context) {
	CONTRACTNBR := c.Param("CONTRACTNBR")
	fmt.Println(CONTRACTNBR)
	trackingtemp = dataservice.GetTracking_SeqFromDb(nil)
	// Parse the request body to get the updated album data
	var updatedTracking_seq models.Tracking_seq
	if err := c.ShouldBindJSON(&updatedTracking_seq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for index, a := range trackingtemp {
		if a.CONTRACTNBR == CONTRACTNBR {
			fmt.Println("In for CONTRACTNBR is:", a.CONTRACTNBR)
			//db := dataservice.ConnectToDb()
			dataservice.UpdateTracking_SeqInDb(updatedTracking_seq)
			// Update the album's data with the new data from the request body
			trackingtemp[index].CONTRACTNBR = updatedTracking_seq.CONTRACTNBR
			trackingtemp[index].SUBSEQ = updatedTracking_seq.SUBSEQ
			trackingtemp[index].GROUPSEQ = updatedTracking_seq.GROUPSEQ
			trackingtemp[index].SEARCHCODE = updatedTracking_seq.SEARCHCODE
			trackingtemp[index].REPTTYPE = updatedTracking_seq.REPTTYPE
			trackingtemp[index].HOMEPLAN = updatedTracking_seq.HOMEPLAN
			trackingtemp[index].MATRLDIST = updatedTracking_seq.MATRLDIST
			trackingtemp[index].DISPCODE = updatedTracking_seq.DISPCODE
			trackingtemp[index].SENTFLAG = updatedTracking_seq.SENTFLAG
			trackingtemp[index].REQ_DATE = updatedTracking_seq.REQ_DATE
			trackingtemp[index].SETUPNAME = updatedTracking_seq.SETUPNAME
			trackingtemp[index].REQ_DATE = updatedTracking_seq.REQ_DATE
			trackingtemp[index].SEARCHCODE = updatedTracking_seq.SEARCHCODE
			trackingtemp[index].REPTTYPE = updatedTracking_seq.REPTTYPE

			c.IndentedJSON(http.StatusOK, trackingtemp[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "trackingtemp not found"})
}
