package api

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/wrighbr/resume-api/models"
)

var experience []models.Experience

func createExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var jsonData models.Experience

	json.NewDecoder(r.Body).Decode(&jsonData)

	jsonData.ID = rand.Intn(1000000)
	experience = append(experience, jsonData)
	json.NewEncoder(w).Encode(&jsonData)

}

func listExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(experience)
}
