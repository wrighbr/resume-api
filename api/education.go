package api

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/wrighbr/resume-api/client"
	"github.com/wrighbr/resume-api/models"
	"google.golang.org/api/iterator"
)

var education []models.Education
var colEducation = "education"

func fillInEducation(r *http.Request) *models.Education {
	var jsonData models.Education
	json.NewDecoder(r.Body).Decode(&jsonData)

	return &models.Education{
		ID:          rand.Intn(1000000),
		StartDate:   jsonData.StartDate,
		EndDate:     jsonData.EndDate,
		Course:      jsonData.Course,
		Institution: jsonData.Institution,
	}
}

// createEducation godoc
// @Summary Creates Education information
// @Description Get the contact information
// @Tags Education
// @Accept  json
// @Produce  json
// @Success 200
// @Security BasicAuth
// @Router /education [post]
func createEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	education := fillInEducation(r)

	client.CreateDocument(colEducation, education)

	location := "/edueducation/" + strconv.Itoa(education.ID)
	w.Header().Set("location", location)
}

// readEducation godoc
// @Summary reads Education information
// @Description Get the contact information
// @Tags Education
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404  "education id not found"
// @Router /education/{id} [get]
func readEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}
	_, jsonData := client.ReadDocument(colEducation, id)

	json.NewEncoder(w).Encode(jsonData)
}

// updateEducation godoc
// @Summary updates Educational information
// @Description Updates the education information
// @Tags Education
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Security BasicAuth
// @Router /education/{id} [put]
func updateEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	jsonData := fillInEducation(r)
	docID, _ := client.ReadDocument(colEducation, id)

	client.UpdateDocument(colEducation, docID, jsonData)
}

// deleteEducation godoc
// @Summary deletes education information
// @Description deletes the eduation information
// @Tags Education
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 404
// @Security BasicAuth
// @Router /education/{id} [delete]
func deleteEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	docID, _ := client.ReadDocument(colEducation, id)

	client.DeleteDocument(colEducation, docID)
}

// getAllEducation godoc
// @Summary Gets all Educational information
// @Description Gets all Educational information
// @Tags Education
// @Accept  json
// @Produce  json
// @Success 200
// @Router /education [get]
func getAllEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	education = nil
	var data models.Education

	ctx := context.Background()
	client := createClient(ctx)

	iter := client.Collection(colEducation).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		err = mapstructure.Decode(doc.Data(), &data)
		if err != nil {
			fmt.Println(err)
		}
		education = append(education, data)
	}

	if jsonResponse {
		json.NewEncoder(w).Encode(education)
	}
}
