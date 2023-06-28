package main

import (
	"github.com/ghostdev-labs/GoUrlShortener/models"
	"github.com/ghostdev-labs/GoUrlShortener/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new router
	router := gin.Default()

	// Setup the router
	routes.SetupRouter(router)

	// Defer the closing of the database connection
	defer models.CloseDB()

	// Run the server
	router.Run()
}
