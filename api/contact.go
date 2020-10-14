package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wrighbr/resume-api/client"
	"github.com/wrighbr/resume-api/models"
)

var collectionContactInfo = "contactInfo"
var contactInfo models.ContactInfo

func getContactInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, doc := client.ReadDocument(collectionContactInfo, 1)

	fmt.Println(doc)
	json.NewEncoder(w).Encode(doc)
}

func createContantInfo(w http.ResponseWriter, r *http.Request) {

	contactInfo := fillinContactInfo(r)

	client.CreateDocument(collectionContactInfo, contactInfo)
}

func updateContantInfo(w http.ResponseWriter, r *http.Request) {

	contactInfo := fillinContactInfo(r)
	id, _ := client.ReadDocument(collectionContactInfo, 1)
	// fmt.Println(id)
	client.UpdateDocument(collectionContactInfo, id, contactInfo)
}

func deleteContantInfo(w http.ResponseWriter, r *http.Request) {
	id, _ := client.ReadDocument(collectionContactInfo, 1)

	client.DeleteDocument(collectionContactInfo, id)

}

func fillinContactInfo(r *http.Request) *models.ContactInfo {

	var jsonData models.ContactInfo
	json.NewDecoder(r.Body).Decode(&jsonData)

	return &models.ContactInfo{
		ID:      1,
		Name:    jsonData.Name,
		Email:   jsonData.Email,
		Mobile:  jsonData.Mobile,
		Address: jsonData.Address,
		Town:    jsonData.Town,
		Country: jsonData.Country,
		Github:  jsonData.Github,
		Website: jsonData.Website,
	}
}
