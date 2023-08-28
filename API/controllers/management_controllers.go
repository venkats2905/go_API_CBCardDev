package controller

import (
	"fmt"
	"go_API_CBCardDev/dataservice"
	"go_API_CBCardDev/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var management []models.Management

func GetManagement(c *gin.Context) {
	fmt.Println("To get all the management present")
	management = dataservice.GetManagementFromDb(nil)
	c.IndentedJSON(http.StatusOK, management)
}

func GetMangementByID(c *gin.Context) {
	management = dataservice.GetManagementFromDb(nil)

	stockid := c.Param("stockcode")
	fmt.Println(stockid)

	for _, a := range management {
		fmt.Println("I'm stockID", stockid, "I'm a.Stockcode", a.StockCode)

		stockCode, _ := strconv.Atoi(stockid) // Convert stockid to an integer
		if a.StockCode == stockCode {
			managementByStockCode, err := dataservice.GetManagementBystockcodeFromDb(stockid)
			if err != nil {
				fmt.Println("Failed to get data:", err)
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve management data"})
				return
			}
			c.IndentedJSON(http.StatusOK, managementByStockCode)
			return
		}

	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "management not found"})
}

func AddManagement(c *gin.Context) {
	fmt.Println("IN POST REQ FUNCTION")
	var newmanagement models.Management

	if err := c.BindJSON(&newmanagement); err != nil {
		return
	}

	dataservice.PostAddManagementToDb(newmanagement)
	management = append(management, newmanagement)
	fmt.Println(management)
	c.IndentedJSON(http.StatusCreated, management)
}

func DeletemanagementByID(c *gin.Context) {
	management = dataservice.GetManagementFromDb(nil)
	stockid := c.Param("stockcode")
	fmt.Println(stockid)

	for index, a := range management {
		stockCode, _ := strconv.Atoi(stockid)
		if a.StockCode == stockCode {
			dataservice.DeleteManagementFromDb(stockid)
			management = append(management[:index], management[index+1:]...)
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "management not found"})
}

func UpdatemanagementByID(c *gin.Context) {
	management = dataservice.GetManagementFromDb(nil)
	stockid := c.Param("stockcode")
	fmt.Println(stockid)

	var updatedmanagement models.Management
	if err := c.ShouldBindJSON(&updatedmanagement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, a := range management {
		stockid, _ := strconv.Atoi(stockid)
		if a.StockCode == stockid {
			dataservice.UpdateManagementInDb(updatedmanagement)
			management[index] = updatedmanagement
			c.IndentedJSON(http.StatusOK, management[index])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "management not found"})
}
