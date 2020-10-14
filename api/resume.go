package api

import (
	"encoding/json"
	"net/http"

	"github.com/wrighbr/resume-api/models"
)

// var resume models.Resume

func getResume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resume := models.Resume{
		ContactInfo: contactInfo,
		Experience:  experience,
	}

	json.NewEncoder(w).Encode(resume)
}
