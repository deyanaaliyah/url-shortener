package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define the URL struct
type Url struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"primaryKey"`
	Url          string `json:"url"`
	ShortenedURL string `json:"shortened_url"`
	Title		 string `json:"title"`
}

// Database connection
var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(&Url{}) // Migrate the schema
}

// Fetch website title
func fetchTitle(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Unknown Title"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "Unknown Title"
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "Unknown Title"
	}

	title := strings.TrimSpace(doc.Find("title").Text())
	if title == "" {
		return "No Title Found"
	}

	return title
}

// Generate a unique shortened URL
func generateShortenedURL() string {
	for {
		// Generate 6 random bytes
		b := make([]byte, 6)
		_, err := rand.Read(b)
		if err != nil {
			panic("failed to generate random ID")
		}

		// Convert to hex (URL-safe)
		shortened := hex.EncodeToString(b)

		// Check if the shortened URL already exists
		var count int64
		db.Model(&Url{}).Where("shortened_url = ?", shortened).Count(&count)

		if count == 0 {
			return shortened
		}
	}
}

// Get all URLs
func GetAllUrls(c *gin.Context) {
	var urls []Url

	if err := db.Find(&urls).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URLs"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": urls})
}

// Create a shortened URL with title
func CreateUrl(c *gin.Context) {
	var newURL Url

	if err := c.ShouldBindJSON(&newURL); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}


	// Generate unique shortened URL
	newURL.ShortenedURL = generateShortenedURL()

	// Fetch and store the website title
	newURL.Title = fetchTitle(newURL.Url)

	// Save to database
	if err := db.Create(&newURL).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "URL created successfully!",
		"original_url":  newURL.Url,
		"shortened_url": newURL.ShortenedURL,
		"title":         newURL.Title,
	})
}

// Redirect to the original URL
func RedirectUrl(c *gin.Context) {
	shortenedURL := c.Param("shortenedUrl")

	// Find the URL by shortened version
	var existingUrl Url
	if err := db.Where("shortened_url = ?", shortenedURL).First(&existingUrl).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Redirect to the original URL
	c.Redirect(http.StatusFound, existingUrl.Url)
}

// Soft delete a shortened URL
func SoftDeleteUrl(c *gin.Context) {
	shortenedURL := c.Param("shortenedUrl")

	var existingUrl Url

	// Check if the URL exists
	if err := db.Where("shortened_url = ?", shortenedURL).First(&existingUrl).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	// Soft delete
	if err := db.Delete(&existingUrl).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't delete URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "URL has been soft deleted successfully!"})
}
