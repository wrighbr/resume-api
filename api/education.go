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

func createEducation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	education := fillInEducation(r)

	client.CreateDocument(colEducation, education)

}

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
