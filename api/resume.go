package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mitchellh/mapstructure"
	"github.com/wrighbr/resume-api/client"
	"github.com/wrighbr/resume-api/models"
)

var jsonResponse = true

func getResume(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse = false

	var contact models.ContactInfo
	_, data := client.ReadDocument(collectionContactInfo, 1)

	err := mapstructure.Decode(data, &contact)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(contact)

	getAllExperience(w, r)
	getAllEducation(w, r)

	resume := models.Resume{
		ContactInfo: contact,
		Experience:  experience,
		Education:   education,
	}

	jsonResponse = true

	json.NewEncoder(w).Encode(resume)
}
