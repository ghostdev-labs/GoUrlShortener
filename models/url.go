package models

import (
	"math/rand"
	"net/url"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// db is the database connection
var db *gorm.DB

// URL is the model for storing shortened URLs
type URL struct {
	gorm.Model
	LongURL string `json:"long_url" gorm:"unique;not null"`
	ShortURL string `json:"short_url" gorm:"unique;not null"`
	AccessCount uint `json:"access_count"`
	LastAccessedAt *time.Time `json:"last_accessed_at"`
	AccessedFromIP string `json:"accessed_from_ip"`
}

// GenerateShortURL generates a short URL for the given long URL
func (url *URL) GenerateShortURL() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	shortURLRunes := make([]rune, 6)
	for i := range shortURLRunes {
		shortURLRunes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	url.ShortURL = string(shortURLRunes)
}

// ParseRequestURI parses the long URL using the net/url package
func (u *URL) ParseRequestURI() (*url.URL, error) {
    return url.ParseRequestURI(u.LongURL)
}

// CreateURL creates a new URL record in the database
func CreateURL(url *URL) error {
	return db.Create(url).Error
}

// GetURLByShortURL retrieves a URL record from the database by short URL
func GetURLByShortURL(shortURL string) (*URL, error) {
	var url URL
	if err := db.Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

// UpdateURL updates a URL record in the database
func UpdateURL(url *URL) error {
	result := db.Save(url)
	return result.Error
}

// GetURLs retrieves all URL records from the database
func GetURLs() ([]*URL, error) {
	var urls []*URL
	if err := db.Find(&urls).Error; err != nil {
		return nil, err
	}
	return urls, nil
}

// DeleteURL deletes a URL record from the database
func DeleteURL(url *URL) error {
	result := db.Delete(url)
	return result.Error
}

// CloseDB closes the database connection
func CloseDB() {
	db.Close()
}
