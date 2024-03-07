package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BrunoPolaski/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
}

func GetPersonalities(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalities)
}
