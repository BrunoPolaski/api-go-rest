package main

import (
	"fmt"

	"github.com/BrunoPolaski/api-go-rest/routes"
)

func main() {
	fmt.Println("Initialising the server...")
	routes.HandleRequest()
}
