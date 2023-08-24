package controller

import (
	"db/dataservice"
	"db/models"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	// _ "github.com/gin-gonic/gin"
)

func GetCard(c *gin.Context) {
	var cards []models.Carddetails
	fmt.Println("To get all the cards present")
	// db := dataservice.ConnectToDb()
	// fmt.Println("\n in getProduction", db)
	cards = dataservice.GetCardDataFromDb(cards)
	c.IndentedJSON(http.StatusOK, cards)
	cards = make([]models.Carddetails, 0)
}

func GetCardByID(c *gin.Context) {
	var cards []models.Carddetails
	cardrevisionid := c.Param("cardrevisionid")
	Cardrevisionid1, err := strconv.Atoi(cardrevisionid)
	fmt.Println(Cardrevisionid1, err, reflect.TypeOf(Cardrevisionid1))
	if err != nil {
		// Handle error (conversion failed)
		// For example, you might return an error response to the client
		// fmt.Println("Else P=====================")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cardrevisionid"})
		return
	}
	//fmt.Println(cardrevisionid)

	// You're looping over productioncards, but you haven't shown where productioncards is defined.
	// Assuming productioncards is a slice of models.carddetails.
	cards = dataservice.GetCardDataFromDb(cards)
	// fmt.Println("-------------------", productioncards)
	for _, a := range cards {
		// fmt.Println("1.--> ", a.Cardrevisionid, cardrevisionid)
		if a.Cardrevisionid == Cardrevisionid1 {
			// fmt.Println("2.--> ", a.Cardrevisionid, Cardrevisionid1)
			// Connect to the database and fetch the production card by contractnbr
			productioncardbycardrevisionid, err := dataservice.GetCardDataByCardrevisionidFromDb(Cardrevisionid1)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve production card data"})
				return
			}
			// Return the fetched production card
			c.IndentedJSON(http.StatusOK, productioncardbycardrevisionid)
			return
		}
	}

	// If no matching production card was found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "card not found"})
}

func AddCard(c *gin.Context) {
	fmt.Println("--------------------------------IN POST REQ FUNCTION----------------------")
	var cards []models.Carddetails
	var newcards models.Carddetails

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newcards); err != nil {
		// fmt.Println("ERROR in getting data")
		return
	}
	// Add the new album to the slice.
	fmt.Println(newcards)
	//db := dataservice.ConnectToDb()
	cards = dataservice.GetCardDataFromDb(cards)
	dataservice.PostAddCardToDb(newcards)
	cards = append(cards, newcards)
	fmt.Println(cards)
	c.IndentedJSON(http.StatusCreated, newcards)
}

func UpdateCardByID(c *gin.Context) {
	var cards []models.Carddetails
	cardrevisionid := c.Param("cardrevisionid")
	Cardrevisionid1, err := strconv.Atoi(cardrevisionid)
	fmt.Println(Cardrevisionid1, err, reflect.TypeOf(Cardrevisionid1))
	fmt.Println(cardrevisionid)
	// Parse the request body to get the updated album data
	var updatedProductioncard models.Carddetails
	if err := c.ShouldBindJSON(&updatedProductioncard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cards = dataservice.GetCardDataFromDb(cards)
	for index, a := range cards {
		if a.Cardrevisionid == Cardrevisionid1 {
			fmt.Println("In for cardrevisionid is:", a.Cardrevisionid)
			//db := dataservice.ConnectToDb()
			dataservice.UpdateCardInDb(updatedProductioncard)
			// Update the album's data with the new data from the request body
			cards[index].Cardrevisionid = updatedProductioncard.Cardrevisionid
			cards[index].Cardtemplatecode = updatedProductioncard.Cardtemplatecode
			cards[index].Comments = updatedProductioncard.Comments
			cards[index].Stockcode = updatedProductioncard.Stockcode
			cards[index].Effectivedate = updatedProductioncard.Effectivedate
			cards[index].Carriertemplatecode = updatedProductioncard.Carriertemplatecode
			cards[index].Enddate = updatedProductioncard.Enddate
			cards[index].Status = updatedProductioncard.Status
			cards[index].Statusdate = updatedProductioncard.Statusdate
			cards[index].Creationdate = updatedProductioncard.Creationdate
			cards[index].Creationuserid = updatedProductioncard.Creationuserid
			cards[index].Lastupdates = updatedProductioncard.Lastupdates
			cards[index].Lastupdateuserid = updatedProductioncard.Lastupdateuserid

			c.IndentedJSON(http.StatusOK, cards[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cards not found"})
}

func DeleteCardByID(c *gin.Context) {
	var cards []models.Carddetails
	cards = dataservice.GetCardDataFromDb(cards)
	cardrevisionid := c.Param("cardrevisionid")
	Cardrevisionid1, err := strconv.Atoi(cardrevisionid)
	fmt.Println(Cardrevisionid1, err, reflect.TypeOf(Cardrevisionid1))
	//fmt.Println(cardrevisionid)
	for index, a := range cards {
		fmt.Println("In for id is:", a.Cardrevisionid)
		if a.Cardrevisionid == Cardrevisionid1 {
			//db := dataservice.ConnectToDb()
			dataservice.DeleteCardFromDb(a.Cardrevisionid)
			cards = append(cards[:index], cards[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "card not found"})
	cards = make([]models.Carddetails, 0)

}
