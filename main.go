package main

import (
	"fmt"

	"github.com/BrunoPolaski/api-go-rest/database"
	"github.com/BrunoPolaski/api-go-rest/models"
	"github.com/BrunoPolaski/api-go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "Bruno", History: "I am a software engineer"},
		{Id: 2, Name: "John", History: "I am a doctor"},
	}
	database.ConnectToDatabase()
	fmt.Println("Initialising the server...")
	routes.HandleRequest()
}
