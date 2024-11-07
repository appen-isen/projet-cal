package main

import (
	"bufio"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
	"time"
)

type LinkAndDate struct {
	Link string `json:"link"`
	Date string `json:"date"`
}

func getDayAndLink(c *gin.Context) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	fileName := os.Getenv("FILE_NAME")
	if fileName == "" {
		log.Fatalf("FILE_NAME not set in .env")
	}

	date := time.Now().Format("02-01-2006")

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing file: %v", err)
		}
	}(file)

	links := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ";")
		if len(parts) == 2 {
			links[parts[0]] = parts[1]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	link, exists := links[date]
	if !exists {
		link = "https://instagram.com/appen_isen"
	}

	c.JSON(200, LinkAndDate{Link: link, Date: date})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Serve static files from the frontend directory
	router.Static("/static", "./frontend")

	// API route
	router.GET("/api/get_day", getDayAndLink)

	// Serve the index.html file
	router.NoRoute(func(c *gin.Context) {
		c.File("../frontend/index.html")
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
