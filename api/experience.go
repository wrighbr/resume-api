package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/wrighbr/resume-api/client"
	"github.com/wrighbr/resume-api/models"
	"google.golang.org/api/iterator"
)

var experience []models.Experience
var colExperience = "experience"

func fillInxperience(r *http.Request) *models.Experience {
	var jsonData models.Experience
	json.NewDecoder(r.Body).Decode(&jsonData)

	return &models.Experience{
		ID:          rand.Intn(1000000),
		StartDate:   jsonData.StartDate,
		EndDate:     jsonData.EndDate,
		Company:     jsonData.Company,
		Role:        jsonData.Role,
		Description: jsonData.Description,
		Location:    jsonData.Location,
	}
}

func createExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	experience := fillInxperience(r)

	client.CreateDocument(colExperience, experience)

}

func readExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}
	_, jsonData := client.ReadDocument(colExperience, id)

	json.NewEncoder(w).Encode(jsonData)
}

func updateExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	jsonData := fillInxperience(r)
	docID, _ := client.ReadDocument(colExperience, id)

	client.UpdateDocument(colExperience, docID, jsonData)
}

func deleteExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params["ID"])
	id, err := strconv.Atoi(params["ID"])
	if err != nil {
		fmt.Println(err)
	}

	docID, _ := client.ReadDocument(colExperience, id)

	client.DeleteDocument(colExperience, docID)
}

func getAllExperience(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var experience []models.Experience
	experience = nil
	var data models.Experience

	ctx := context.Background()
	client := createClient(ctx)

	iter := client.Collection(colExperience).Documents(ctx)
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
		experience = append(experience, data)
	}
	// jsonResponse = true
	if jsonResponse {
		json.NewEncoder(w).Encode(experience)
	}
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GCP_PROJECT_ID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}
