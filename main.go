package main

import (
	"fmt"

	"github.com/BrunoPolaski/api-go-rest/models"
	"github.com/BrunoPolaski/api-go-rest/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Bruno", History: "I am a software engineer"},
		{Name: "John", History: "I am a doctor"},
	}
	fmt.Println("Initialising the server...")
	routes.HandleRequest()
}
