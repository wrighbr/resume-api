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

// getContctInfo godoc
// @Summary Gets the contact information
// @Description Get the contact information
// @Tags contact
// @Accept  json
// @Produce  json
// @Success 200
// @Router /contactinfo [get]
func getContactInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, doc := client.ReadDocument(collectionContactInfo, 1)

	fmt.Println(doc)
	json.NewEncoder(w).Encode(doc)
}

// createContctInfo godoc
// @Summary Creates contact information
// @Param name body models.ContactInfo true "Personal Contact Information"
// @Tags contact
// @Accept  json
// @Success 200
// @Router /contactinfo [post]
func createContantInfo(w http.ResponseWriter, r *http.Request) {

	contactInfo := fillinContactInfo(r)

	client.CreateDocument(collectionContactInfo, contactInfo)
}

// ContctInfo godoc
// @Summary Updates contact information
// @Tags contact
// @Accept  json
// @Success 200
// @Router /contactinfo [put]
func updateContantInfo(w http.ResponseWriter, r *http.Request) {

	contactInfo := fillinContactInfo(r)
	id, _ := client.ReadDocument(collectionContactInfo, 1)
	// fmt.Println(id)
	client.UpdateDocument(collectionContactInfo, id, contactInfo)
}

// createContctInfo godoc
// @Summary Deletes contact information
// @Tags contact
// @Accept  json
// @Success 200
// @Router /contactinfo [delete]
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
