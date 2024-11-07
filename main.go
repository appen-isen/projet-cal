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
	//Charge le fichier .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//Récupération du chemin du fichier contenant les dates et les liens depuis le .env
	fileName := os.Getenv("FILE_NAME")
	if fileName == "" {
		log.Fatalf("FILE_NAME not set in .env")
	}

	//Récupération de la date du jour
	date := time.Now().Format("02-01-2006")

	//Lis le fichier de contenant les dates et les liens (date_link.txt par exemple) à son chemin spécifié dans le .env
	//Si la date du jour est présente dans le fichier, renvoie le lien associé
	//Le fichier doit être sous la forme suivante :
	//01-12-2024;https://instagram.com/appen_isen
	//Sinon, renvoie un lien par défaut
	//Le lien par défault est https://instagram.com/appen_isen
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	//Crée une carte pour stocker les dates et les liens
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

	//Cherche le lien associé à la date du jour
	link, exists := links[date]
	if !exists {
		link = "https://instagram.com/appen_isen"
	}

	c.JSON(200, LinkAndDate{Link: link, Date: date})
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//Définition de la route /get_day avec la méthode GET et la fonction getDayAndLink qui renvoie le jour sous la forme d'un string (14-12-2024) et un lien sous la forme d'un string (https://www.google.com)
	router.GET("/api/get_day", getDayAndLink)

	router.Run(":8080")
}
