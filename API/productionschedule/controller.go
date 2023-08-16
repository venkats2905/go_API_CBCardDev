package productionschedule

import (
	"db/dataservice"
	"db/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Productionschedules []models.Productionschedule

func GetProductionschedule(c *gin.Context) {
	fmt.Println("To get all the Productionschedules present")
	db := dataservice.ConnectToDb()
	fmt.Println("\n in getProduction", db)
	Productionschedules = dataservice.GetProductionScheduleFromDb(Productionschedules)
	c.IndentedJSON(http.StatusOK, Productionschedules)
	Productionschedules = make([]models.Productionschedule, 0)
}

func GetProductionscheduleByID(c *gin.Context) {
	jobname := c.Param("jobname")
	fmt.Println(jobname)
	fmt.Println("production", Productionschedules)
	Productionschedules = dataservice.GetProductionScheduleFromDb(Productionschedules)

	// You're looping over Productionschedules, but you haven't shown where Productionschedules is defined.
	// Assuming Productionschedules is a slice of models.Productioncard.
	for _, a := range Productionschedules {
		fmt.Println(a.Jobname, jobname)
		if a.Jobname == jobname {
			fmt.Println(a.Jobname, jobname)
			// Connect to the database and fetch the Production schedule by jobname
			Productionschedulebyjobname, err := dataservice.GetProductionscheduleByJobnameFromDb(jobname)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve Production schedule data"})
				return
			}
			// Return the fetched Production schedule
			c.IndentedJSON(http.StatusOK, Productionschedulebyjobname)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Production schedule not found"})

}

func AddProductionschedule(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var newproductionschedules models.Productionschedule

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newproductionschedules); err != nil {
		// fmt.Println("ERROR in getting data")
		return
	}
	// Add the new album to the slice.
	fmt.Println(newproductionschedules)
	Productionschedules = dataservice.GetProductionScheduleFromDb(Productionschedules)
	dataservice.PostAddProductionscheduleToDb(newproductionschedules)
	Productionschedules = append(Productionschedules, newproductionschedules)
	fmt.Println(Productionschedules)
	c.IndentedJSON(http.StatusCreated, newproductionschedules)
}

func DeleteProductionscheduleByID(c *gin.Context) {
	jobname := c.Param("jobname")
	fmt.Println(jobname)
	Productionschedules = dataservice.GetProductionScheduleFromDb(Productionschedules)
	for index, a := range Productionschedules {
		fmt.Println("In for id is:", a.Jobname)
		if a.Jobname == jobname {
			//db := dataservice.ConnectToDb()
			dataservice.DeletescheduleAlbumFromDb(a.Jobname)
			Productionschedules = append(Productionschedules[:index], Productionschedules[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			Productionschedules = make([]models.Productionschedule, 0)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "productionschedules not found"})

}

func UpdateProductionscheduleByID(c *gin.Context) {
	jobname := c.Param("jobname")
	fmt.Println(jobname)
	// Parse the request body to get the updated album data
	var updatedProductionschedule models.Productionschedule
	if err := c.ShouldBindJSON(&updatedProductionschedule); err != nil {
		fmt.Println("Getting error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Productionschedules = dataservice.GetProductionScheduleFromDb(Productionschedules)
	for index, a := range Productionschedules {
		if a.Jobname == jobname {
			fmt.Println("In for jobname is:", a.Jobname)
			dataservice.UpdateProductionscheduleInDb(updatedProductionschedule)
			// Update the album's data with the new data from the request body
			Productionschedules[index].Proddate = updatedProductionschedule.Proddate
			Productionschedules[index].Jobname = updatedProductionschedule.Jobname
			Productionschedules[index].Jobqueue = updatedProductionschedule.Jobqueue
			Productionschedules[index].Fullpathname = updatedProductionschedule.Fullpathname
			Productionschedules[index].Cardcount = updatedProductionschedule.Cardcount
			Productionschedules[index].Carriercount = updatedProductionschedule.Carriercount
			Productionschedules[index].Requestdate = updatedProductionschedule.Requestdate
			Productionschedules[index].Sentdate = updatedProductionschedule.Sentdate
			Productionschedules[index].Completedate = updatedProductionschedule.Completedate
			Productionschedules[index].Issues = updatedProductionschedule.Issues
			Productionschedules[index].Status = updatedProductionschedule.Status
			Productionschedules[index].Filesequence = updatedProductionschedule.Filesequence
			Productionschedules[index].Completedby = updatedProductionschedule.Completedby
			Productionschedules[index].Veridiedby = updatedProductionschedule.Veridiedby
			Productionschedules[index].Notes = updatedProductionschedule.Notes
			Productionschedules[index].Embossedcards = updatedProductionschedule.Embossedcards
			Productionschedules[index].Embossedcarriers = updatedProductionschedule.Embossedcarriers
			Productionschedules[index].Heldcards = updatedProductionschedule.Heldcards
			Productionschedules[index].Heldcarriers = updatedProductionschedule.Heldcarriers

			c.IndentedJSON(http.StatusOK, Productionschedules[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "productionschedules not found"})
}
