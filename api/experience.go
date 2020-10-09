package api

import (
	"encoding/json"
	"net/http"

	"github.com/wrighbr/resume-api/models"
)

var experience models.Experience

func createExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var jsonData models.Experience
	json.NewDecoder(r.Body).Decode(&jsonData)

}
