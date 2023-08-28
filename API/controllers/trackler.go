package controller

import (
	"fmt"
	"go_API_CBCardDev/dataservice"
	dbase "go_API_CBCardDev/dataservice"
	"go_API_CBCardDev/models"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Trackrecord []models.Trackingsequence_archive

func Gettrackler(c *gin.Context) {
	fmt.Println("To get all the productioncards present")
	db := dbase.ConnectToDb()
	fmt.Println("\n in getProduction", db)
	Trackrecord = dbase.GetTracking(Trackrecord, db)
	c.IndentedJSON(http.StatusOK, Trackrecord)
	Trackrecord = make([]models.Trackingsequence_archive, 0)
}

func GetBySID(c *gin.Context) {
	SID := c.Param("SID")
	S, err := strconv.Atoi(SID)
	fmt.Println(S, err, reflect.TypeOf(S))
	db := dataservice.ConnectToDb()
	Trackrecord = dataservice.GetTracking(Trackrecord, db)
	for _, a := range Trackrecord {
		fmt.Println(a.SID, S)

		if a.SID == S {
			fmt.Println(a.SID, S)
			productioncardbycontractnbr, err := dataservice.TrackbySID(S, db)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve production card data"})
				return
			}

			c.IndentedJSON(http.StatusOK, productioncardbycontractnbr)
			Trackrecord = make([]models.Trackingsequence_archive, 0)
			return
		}
	}

}

func Adddata(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var newTrack models.Trackingsequence_archive

	if err := c.BindJSON(&newTrack); err != nil {
		return
	}
	fmt.Println(newTrack)
	db := dataservice.ConnectToDb()
	Trackrecord = dataservice.GetTracking(Trackrecord, db)
	dataservice.Postrecord(newTrack, db)
	Trackrecord = append(Trackrecord, newTrack)
	fmt.Println(Trackrecord)
	c.IndentedJSON(http.StatusCreated, Trackrecord)
	Trackrecord = make([]models.Trackingsequence_archive, 0)
}

func UpdatetracklerbyID(c *gin.Context) {
	fmt.Println("------UPDADTE-------------")
	SID := c.Param("SID")
	S, err := strconv.Atoi(SID)
	fmt.Println(S, err, reflect.TypeOf(S))
	var updateTrackingsequence_archive models.Trackingsequence_archive
	if err := c.ShouldBindJSON(&updateTrackingsequence_archive); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := dataservice.ConnectToDb()
	Trackrecord = dataservice.GetTracking(Trackrecord, db)
	for index, a := range Trackrecord {
		if a.SID == S {
			fmt.Println("In for Contractnbr is:", a.SID)
			dataservice.Updaterecord(updateTrackingsequence_archive, db, SID)
			Trackrecord[index].SID = updateTrackingsequence_archive.SID
			Trackrecord[index].CONTRACTNBR = updateTrackingsequence_archive.CONTRACTNBR
			Trackrecord[index].SUBSEQ = updateTrackingsequence_archive.SUBSEQ
			Trackrecord[index].GROUPSEQ = updateTrackingsequence_archive.GROUPSEQ
			Trackrecord[index].GROUPNBR = updateTrackingsequence_archive.GROUPNBR
			Trackrecord[index].SUFFIXNBR = updateTrackingsequence_archive.SUFFIXNBR
			Trackrecord[index].HOMEPLAN = updateTrackingsequence_archive.HOMEPLAN
			Trackrecord[index].IMB = updateTrackingsequence_archive.IMB
			Trackrecord[index].MATRLDIST = updateTrackingsequence_archive.MATRLDIST
			Trackrecord[index].DISPCODE = updateTrackingsequence_archive.DISPCODE
			Trackrecord[index].MAILDATE = updateTrackingsequence_archive.MAILDATE
			Trackrecord[index].SENTFLAG = updateTrackingsequence_archive.SENTFLAG
			Trackrecord[index].CARDTEMPLATECODE = updateTrackingsequence_archive.CARDTEMPLATECODE
			Trackrecord[index].SETUPNAME = updateTrackingsequence_archive.SETUPNAME
			Trackrecord[index].REQ_DATE = updateTrackingsequence_archive.REQ_DATE
			Trackrecord[index].SEARCHCODE = updateTrackingsequence_archive.SEARCHCODE
			Trackrecord[index].RELEASE_DT = updateTrackingsequence_archive.RELEASE_DT
			Trackrecord[index].REPTTYPE = updateTrackingsequence_archive.REPTTYPE
			Trackrecord[index].FEPFLAG = updateTrackingsequence_archive.FEPFLAG

			c.IndentedJSON(http.StatusOK, Trackrecord[index])
			Trackrecord = make([]models.Trackingsequence_archive, 0)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "productioncards not found"})
}

func DeletetracklerByID(c *gin.Context) {
	db := dataservice.ConnectToDb()
	Trackrecord = dataservice.GetTracking(Trackrecord, db)
	SID := c.Param("SID")
	S, err := strconv.Atoi(SID)
	fmt.Println(S, err, reflect.TypeOf(S))

	for index, a := range Trackrecord {
		fmt.Println("In for id is:", a.SID)
		if a.SID == S {
			dataservice.DeleteSID(a.SID, db)
			Trackrecord = append(Trackrecord[:index], Trackrecord[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			Trackrecord = make([]models.Trackingsequence_archive, 0)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "productioncard not found"})

}
