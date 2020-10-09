package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wrighbr/resume-api/models"
)

var contactInfo models.ContactInfo

func getContactInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contactInfo)
}

func updateContantInfo(w http.ResponseWriter, r *http.Request) {
	var jsonData models.ContactInfo
	json.NewDecoder(r.Body).Decode(&jsonData)

	contactInfo = models.ContactInfo{
		Name:    jsonData.Name,
		Email:   jsonData.Email,
		Mobile:  jsonData.Mobile,
		Address: jsonData.Address,
		Town:    jsonData.Town,
		Country: jsonData.Country,
		Github:  jsonData.Github,
		Website: jsonData.Website,
	}

	fmt.Println("Print info here")
	fmt.Println(contactInfo)
	json.NewEncoder(w).Encode(contactInfo)
}

func deleteContantInfo(w http.ResponseWriter, r *http.Request) {
	contactInfo = models.ContactInfo{
		Name:    "",
		Email:   "",
		Mobile:  "",
		Address: "",
		Town:    "",
		Country: "",
		Github:  "",
		Website: "",
	}
	json.NewEncoder(w).Encode(contactInfo)
}
