package controllers

import (
	"time"

	"github.com/ghostdev-labs/GoUrlShortener/models"
	"github.com/gin-gonic/gin"
)

// CreateShortURL creates a new short URL record from the long URL
func CreateShortURL(c *gin.Context) {
	var url models.URL
	if err := c.BindJSON(&url); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Validate the long URL
	if _, err := url.ParseRequestURI(); err != nil {
        c.JSON(400, gin.H{"error": "Invalid URL"})
        return
    }

	// Generate a short URL
	url.GenerateShortURL()

	// Create a new URL record in the database
	if err := models.CreateURL(&url); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"short_url": url.ShortURL})
}

// RedirectShortURL redirects the user to the long URL
func RedirectShortURL(c *gin.Context) {
	shortURL := c.Param("short_url")
	url, err := models.GetURLByShortURL(shortURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "URL not found"})
		return
	}

	// Update the URL record
	url.AccessCount++
	now := time.Now()
	url.LastAccessedAt = &now
	url.AccessedFromIP = c.ClientIP()
	if err := models.UpdateURL(url); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(302, url.LongURL)
}

// GetURLStats retrieves the stats for a short URL
func GetURLStats(c *gin.Context) {
	shortURL := c.Param("short_url")
	url, err := models.GetURLByShortURL(shortURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "URL not found"})
		return
	}

	c.JSON(200, gin.H{
		"long_url": url.LongURL,
		"short_url": url.ShortURL,
		"access_count": url.AccessCount,
		"last_accessed_at": url.LastAccessedAt,
		"accessed_from_ip": url.AccessedFromIP,
	})
}



// GetURLs retrieves all URL records from the database
func GetURLs(c *gin.Context) {
	urls, err := models.GetURLs()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, urls)
}
