package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"ginius/gin-library/data-access/models"
	"ginius/gin-library/data-access/data"
)

func main() {
	router:= gin.Default()
	router.GET("/gins", getGins)
	router.GET("/gins/:id", getGinByID)
	router.POST("/gins", postGin)

	router.Run("localhost:8080")
}

// getGins returns all gin entries
func getGins(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, data.Gins)
}

// postGin adds a new gin entry from the JSON received in the request body
func postGin(c *gin.Context) {
	var newGin models.GinEntry

	if err := c.BindJSON(&newGin); err != nil {
		return
	}

	data.Gins = append(data.Gins, newGin)
	c.IndentedJSON(http.StatusCreated, newGin)
}

// getGinByID returns a gin entry by its ID
func getGinByID(c *gin.Context) {
	id := c.Param("id")
	ginID, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid gin ID"})
		return
	}

	for _, gin := range data.Gins {
		if gin.GinID == ginID {
			c.IndentedJSON(http.StatusOK, gin)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "gin not found"})
}
