package router

import (
	"ginius/gin-library/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the gin router
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/gins", handlers.GetGins)
	router.GET("/gins/:id", handlers.GetGinByID)
	router.POST("/gins", handlers.PostGin)

	return router
}
