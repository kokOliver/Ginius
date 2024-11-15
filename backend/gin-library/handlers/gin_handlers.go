package handlers

import (
	"net/http"
	"strconv"

	"ginius/gin-library/data"
	"ginius/gin-library/models"
	"github.com/gin-gonic/gin"
)

// GetGins returns all gin entries
func GetGins(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Gins)
}

// PostGin adds a new gin entry from the JSON received in the request body
func PostGin(c *gin.Context) {
	var newGin models.GinEntry

	if err := c.BindJSON(&newGin); err != nil {
		return
	}

	data.Gins = append(data.Gins, newGin)
	c.IndentedJSON(http.StatusCreated, newGin)
}

// GetGinByID returns a gin entry by its ID
func GetGinByID(c *gin.Context) {
	id := c.Param("id")
	ginID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid gin ID"})
		return
	}

	for _, ginEntry := range data.Gins {
		if ginEntry.GinID == ginID {
			c.IndentedJSON(http.StatusOK, ginEntry)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "gin not found"})
}
