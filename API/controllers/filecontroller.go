package controller

import (
	"fmt"
	"go_API_CBCardDev/dataservice"
	"go_API_CBCardDev/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var processedfiles []models.ProcessedFile

func Getprocessedfiles(c *gin.Context) {
	fmt.Println("To get all the Processedfile present")
	processedfiles = dataservice.GetProcessedfilesFromDb(nil)
	c.IndentedJSON(http.StatusOK, processedfiles)
}

// GET THE DATA REQUESTED FOR THE SPECIFIC ID
func GetprocessedfileByID(c *gin.Context) {
	fileid := c.Param("fileid")
	fmt.Println(fileid, processedfiles)

	// You're looping over processedfiles, but you haven't shown where processedfiles is defined.
	// Assuming processedfiles is a slice of models.Productioncard.
	processedfiles = dataservice.GetProcessedfilesFromDb(nil)
	for _, a := range processedfiles {
		fmt.Println(a.Fileid, fileid)
		fileid, _ := strconv.Atoi(fileid)
		if a.Fileid == fileid {
			fmt.Println(a.Fileid, fileid)
			// Connect to the database and fetch the PROCESSED FILE by FILEID
			processedfilebyfileid, err := dataservice.GetprocessedfilesByfileidFromDb(fileid)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve PROCESSED FILE ID  data"})
				return
			}
			// Return the fetched production card
			c.IndentedJSON(http.StatusOK, processedfilebyfileid)
			return
		}
	}
	// If no matching PROCESSED FILE was found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "PROCESSED FILE not found"})
}

func Addprocessedfile(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var newprocessedfiles models.ProcessedFile

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newprocessedfiles); err != nil {
		// fmt.Println("ERROR in getting data")
		return
	}
	// Add the new album to the slice.
	fmt.Println(newprocessedfiles)
	//db := dataservice.ConnectToDb()
	processedfiles = dataservice.GetProcessedfilesFromDb(processedfiles)
	dataservice.PostAddProcessedfileToDb(newprocessedfiles)
	processedfiles = append(processedfiles, newprocessedfiles)
	fmt.Println(processedfiles)
	c.IndentedJSON(http.StatusCreated, newprocessedfiles)
}

func DeleteprocessedfileByID(c *gin.Context) {
	processedfiles = dataservice.GetProcessedfilesFromDb(processedfiles)
	fileid := c.Param("fileid")
	fmt.Println(fileid)
	for index, a := range processedfiles {
		fmt.Println("In for id is:", a.Fileid)
		if strconv.Itoa(a.Fileid) == fileid {
			fmt.Println(a.Fileid, fileid)
			//db := dataservice.ConnectToDb()
			dataservice.DeleteProcessedfileFromDb(a.Fileid)
			processedfiles = append(processedfiles[:index], processedfiles[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "processedfile not found"})
}

func UpdateprocessedfileByID(c *gin.Context) {
	fileid := c.Param("fileid")
	fmt.Println(fileid)
	// Parse the request body to get the updated album data
	var updateprocessedfile models.ProcessedFile
	if err := c.ShouldBindJSON(&updateprocessedfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	processedfiles = dataservice.GetProcessedfilesFromDb(processedfiles)
	for index, a := range processedfiles {
		if strconv.Itoa(a.Fileid) == fileid {
			fmt.Println(a.Fileid, fileid)
			fmt.Println("In for fileid is:", a.Fileid)
			//db := dataservice.ConnectToDb()
			dataservice.UpdateProcessedfileInDb(updateprocessedfile)
			// Update the album's data with the new data from the request body
			processedfiles[index].Fileid = updateprocessedfile.Fileid
			processedfiles[index].Source_system = updateprocessedfile.Source_system
			processedfiles[index].Platform = updateprocessedfile.Platform
			processedfiles[index].Isproduction = updateprocessedfile.Isproduction
			processedfiles[index].Header_date = updateprocessedfile.Header_date
			processedfiles[index].Process_date = updateprocessedfile.Process_date
			processedfiles[index].Start_seq = updateprocessedfile.Start_seq
			processedfiles[index].End_seq = updateprocessedfile.End_seq
			processedfiles[index].File_name = updateprocessedfile.File_name
			processedfiles[index].Mabx_csm_fileid = updateprocessedfile.Mabx_csm_fileid
			c.IndentedJSON(http.StatusOK, processedfiles[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "processedfiles not found"})
}
