package routes

import (
	"github.com/ghostdev-labs/GoUrlShortener/controllers"
	"github.com/ghostdev-labs/GoUrlShortener/models"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router for the application
func SetupRouter(router *gin.Engine) {
	models.InitDB()

	// Create a new short URL
	router.POST("/shorten", controllers.CreateShortURL)

	// Redirect to the long URL
	router.GET("/:short_url", controllers.RedirectShortURL)

	// Get the stats for a short URL
	router.GET("/:short_url/stats", controllers.GetURLStats)

	router.Use(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "URL not found"})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "URL not found"})
	})

}