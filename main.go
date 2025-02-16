package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ip := flag.String("ip", "localhost", "IP address to bind to")
	port := flag.String("port", "8080", "Port to bind to")
	flag.Parse()
	addr := fmt.Sprintf("%s:%s", *ip, *port)
	log.Printf("Starting server on %s", addr)

	router := gin.Default()

	// Configure CORS 
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                             // Allow all origins, you may want to restrict this in production
	config.AllowHeaders = []string{"Authorization", "Content-Type"} // Include the Content-Type header
	router.Use(cors.New(config))

	router.GET("/", GetAllUrls)
	router.GET("/:shortenedUrl", RedirectUrl)
	router.POST("/shorten", CreateUrl)
	router.DELETE("/delete/:shortenedUrl", SoftDeleteUrl)

	router.Run(addr)
}
