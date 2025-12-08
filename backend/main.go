package main

import (
	"bufio"
	"github.com/gin-contrib/cors"
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

	link/*, exists*/ := links[date]
	/*if !exists {
		link = "https://instagram.com/appen_isen"
	}*/

	c.JSON(200, LinkAndDate{Link: link, Date: date})
}

func getLegacyLink(c *gin.Context) {
	requestedDate := c.Query("date")
	currentDate := time.Now().Format("02-01-2006")

	// Si la date demandée n'est pas la date du jour -> FLAG
	if requestedDate != "" && requestedDate != currentDate {
		c.JSON(200, LinkAndDate{
			Link: "CTF{55460e0d0aec5088a008ce54e035e67f}",
			Date: requestedDate,
		})
		return
	}

	// Sinon, comportement normal (on réutilise la logique de getDayAndLink ou on redirige)
	// Pour faire simple ici, on appelle la même logique de récupération
	getDayAndLink(c)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Serve static files from the frontend directory
	router.Static("/static", "../frontend")

	// API route
	router.GET("/api/get_day", getDayAndLink)
	router.GET("/api/get_link", getLegacyLink)

	// Serve the index.html file for all other routes
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/static/") {
			c.File("../frontend" + path)
		} else {
			c.File("../frontend/index.html")
		}
	})

	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
